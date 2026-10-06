package server

// Owner: Member C
//
// Stored chat history, safe for concurrent use.

import "sync"

// History keeps every chat message in the order it was sent.
// The mutex makes it safe when many client goroutines use it at once.
type History struct {
	mu       sync.Mutex
	messages []Message
}

// Add stores a message at the end of the history.
func (h *History) Add(m Message) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.messages = append(h.messages, m)
}

// All returns a COPY of the messages, so callers cannot change the
// history (or race with Add) while they loop over the result.
func (h *History) All() []Message {
	h.mu.Lock()
	defer h.mu.Unlock()
	out := make([]Message, len(h.messages))
	copy(out, h.messages)
	return out
}
