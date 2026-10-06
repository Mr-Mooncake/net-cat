package server

import (
	"fmt"
	"log"
	"net"
	"sync"
)

// Owner: Member A
//
// Listener, accept loop, connection limit.
//
// TODO: see CHECKLIST.txt for your milestones and rules.

const maxClients = 10

var (
	mu sync.Mutex
	activeCount int
)
func Start(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	defer listener.Close()
	fmt.Printf("listening on the port :%s\n", port)
	
	// keeps accepting clients forever
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		if !acquireSlot() {
		_, err := conn.Write([]byte("Chat is full. Try again later.\n"))
		if err != nil {
			log.Printf("Failed to notify rejected client: %v", err)
		}
		conn.Close()
		continue
		}
		go func() {
		defer releaseSlot()
		defer conn.Close()
		handleClient(conn)
		} ()
	}
}

func acquireSlot() bool {
	mu.Lock()
	defer mu.Unlock()
	if activeCount >= maxClients {
		return false
	}
	activeCount++
	return true
}

func releaseSlot() {
	mu.Lock()
	defer mu.Unlock()
	activeCount--
}
