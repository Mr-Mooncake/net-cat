package server

// Owner: Member C
//
// Client registry, broadcast, join/leave notices, history replay.
//
// NOTE: added by Member B only because Milestone 7 needs it. Had to do your part so i can finish milestone 7
// (registry + broadcast + history replay). Member C can replace or merge.

import (
	"sync"
	"time"
)

// writeTimeout stops one slow/broken client from blocking the others.
const writeTimeout = 2 * time.Second

var (
	clientsMu   sync.RWMutex                 // protects the clients map
	clients     = make(map[*Client]struct{}) // everyone currently connected
	chatHistory History                      // all chat messages, in order
)

// registerClient adds a client to the list of connected clients.
func registerClient(c *Client) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	clients[c] = struct{}{}
}

// unregisterClient removes a client (called when they leave).
func unregisterClient(c *Client) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	delete(clients, c)
}

// connectedClients returns a copy of the list so we never write
// to a connection while holding the lock.
func connectedClients() []*Client {
	clientsMu.RLock()
	defer clientsMu.RUnlock()
	list := make([]*Client, 0, len(clients))
	for c := range clients {
		list = append(list, c)
	}
	return list
}

// broadcast sends text to every client except skip (use nil to send to all).
// Each receiver gets a fresh prompt afterwards, like the subject's example.
func broadcast(text string, skip *Client) {
	var wg sync.WaitGroup
	for _, c := range connectedClients() {
		if c == skip {
			continue
		}
		wg.Add(1)
		go func(c *Client) { // one goroutine per client: a stuck one can't block the rest
			defer wg.Done()
			if err := c.send("\n" + text + "\n" + formatPrompt(c.Name)); err != nil {
				c.Conn.Close() // broken client: close it so its own session ends
			}
		}(c)
	}
	wg.Wait()
}

// sendHistory replays all previous messages to a newly joined client.
func sendHistory(c *Client) error {
	for _, m := range chatHistory.All() {
		if err := c.send(m.String() + "\n"); err != nil {
			return err
		}
	}
	return nil
}
