package server

// Owner: Member C
//
// Message formatting: [YYYY-MM-DD HH:MM:SS][name]:message

import (
	"fmt"
	"time"
)

// timeLayout is Go's way of writing YYYY-MM-DD HH:MM:SS.
const timeLayout = "2006-01-02 15:04:05"

// Message is one chat message sent by a client.
type Message struct {
	Time time.Time
	Name string
	Text string
}

// newMessage creates a message stamped with the current time.
func newMessage(name, text string) Message {
	return Message{Time: time.Now(), Name: name, Text: text}
}

// String returns the message as shown in the chat:
// [2020-01-20 15:48:41][name]:text
func (m Message) String() string {
	return fmt.Sprintf("[%s][%s]:%s", m.Time.Format(timeLayout), m.Name, m.Text)
}

// joinNotice and leaveNotice build the server notices (not stored in history).
func joinNotice(name string) string  { return name + " has joined our chat..." }
func leaveNotice(name string) string { return name + " has left our chat..." }
