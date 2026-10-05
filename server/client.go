package server

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

// Owner: Member B
//
// Client type and the per-client session (name, message loop, leaving).

// Client represents a connected chat user with their connection and username.
type Client struct {
	Conn net.Conn
	Name string
}

// readName prompts and reads lines until a non-empty name is entered.
func readName(conn net.Conn, reader *bufio.Reader) (string, error) {
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return "", err // Disconnected or read error
		}

		name := strings.TrimSpace(line)
		if name != "" {
			return name, nil // Valid non-empty name accepted
		}

		// Empty input: reject and ask for the name again
		if _, err := conn.Write([]byte(PromptName)); err != nil {
			return "", err
		}
	}
}

// formatPrompt returns the prompt string: [YYYY-MM-DD HH:MM:SS][name]:
func formatPrompt(name string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s][%s]:", timestamp, name)
}

// showPrompt writes the prompt to the client so they know it's their turn.
func (c *Client) showPrompt() error {
	_, err := c.Conn.Write([]byte(formatPrompt(c.Name)))
	return err
}

// chatLoop reads messages line by line until the client leaves.
func chatLoop(client *Client, reader *bufio.Reader) {
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			// io.EOF = normal leave (Ctrl+C / terminal closed); anything else is logged
			if !errors.Is(err, io.EOF) {
				log.Printf("read error from %s: %v", client.Name, err)
			}
			return
		}

		// Remove spaces and the trailing newline (\n or \r\n)
		msg := strings.TrimSpace(line)

		// "exit" ends the session cleanly
		if msg == "exit" {
			return
		}

		// Drop empty messages: don't send them, just re-show the prompt
		if msg == "" {
			if err := client.showPrompt(); err != nil {
				return
			}
			continue
		}

		// TODO (Member C, Milestone 6): hand msg to broadcast here

		// Re-show the prompt for the next message
		if err := client.showPrompt(); err != nil {
			return
		}
	}
}

// handleClient runs the per-client session: welcome, name, then chat loop.
func handleClient(conn net.Conn) {
	// Always close the connection when the session ends
	defer conn.Close()

	// 1. Send welcome message, Linux logo, and the name prompt
	if _, err := conn.Write([]byte(WelcomeMessage + PromptName)); err != nil {
		return
	}

	// 2. Read and validate the username
	reader := bufio.NewReader(conn)
	name, err := readName(conn, reader)
	if err != nil {
		return
	}

	// 3. Store the name and connection in a Client
	client := &Client{
		Conn: conn,
		Name: name,
	}

	// 4. Show the first message prompt
	if err := client.showPrompt(); err != nil {
		return
	}

	// 5. Chat loop (Milestone 5): runs until the client leaves
	chatLoop(client, reader)
}
