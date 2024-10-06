package chats

import (
	"log"
	"sync"

	"github.com/google/uuid"
)

// Server struct manages multiple chat
// rooms and handles client registration,
// unregistration, and message broadcasting.
type Server struct {
	// Chats is a map where each chat
	// room is identified by its UUID.
	Chats map[uuid.UUID]*Chat

	// Register channel for new
	// clients to join a chat room.
	Register chan *Client

	// Unregister channel for
	// clients leaving a chat room.
	Unregister chan *Client

	// Broadcast channel for sending messages
	// to all clients in a chat room.
	Broadcast chan *Message

	// mutex ensures thread-safe
	// access to the server's state.
	mutex sync.Mutex
}

// NewServer initializes a new Server instance
// with channels and an empty map of chats.
func NewServer() *Server {
	return &Server{
		Chats:      make(map[uuid.UUID]*Chat),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),

		// Buffered channel with
		// capacity of 5 messages
		Broadcast: make(chan *Message, 5),
	}
}

// Run is the main loop that continuously listens for Register, Unregister,
// and Broadcast events, delegating tasks to respective handler functions.
func (s *Server) Run() {
	for {
		select {
		case c := <-s.Register:
			s.register(c)
		case c := <-s.Unregister:
			s.unregister(c)
		case m := <-s.Broadcast:
			s.broadcast(m)
		}
	}
}

// register adds a client to the chat room if the chat exists.
// If the chat doesn't exist, it logs an error and exits the function.
func (s *Server) register(c *Client) {
	// Lock to ensure thread-safe modifications
	s.mutex.Lock()

	defer s.mutex.Unlock()

	chat, exists := s.Chats[c.ChatID]

	if !exists {
		log.Printf("Chat %s does not exist", c.ChatID)
		return
	}

	id := c.Details.ID

	// If the client doesn't already exist in the chat room, add them.
	if _, exists := chat.Clients[id]; !exists {
		chat.Clients[id] = c

		log.Printf("Client %s joined chat %s", c.Details.ID, chat.ID)
	}
}

// unregister removes a client from a chat room. If the chat becomes empty after the client leaves,
// the chat room is deleted. The function also broadcasts a message indicating the client has left.
func (s *Server) unregister(c *Client) {
	s.mutex.Lock()

	defer s.mutex.Unlock()

	chat, exists := s.Chats[c.ChatID]

	if !exists {
		log.Printf("Chat %s does not exist", c.ChatID)
		return
	}

	id := c.Details.ID

	// Ignore if client exists.
	if _, exists := chat.Clients[id]; !exists {
		return
	}

	delete(chat.Clients, id)

	close(c.Message) // Close the client's message channel.

	log.Printf("Client %s left chat %s", id, chat.ID)

	// If no clients are left in the
	// chat, delete the chat room.
	if len(chat.Clients) <= 0 {
		delete(s.Chats, chat.ID)
		log.Printf("Chat %s is now empty, deleting...", chat.ID)
		return
	}

	// Broadcast a message to other clients in
	// the chat room indicating the client left.
	s.Broadcast <- &Message{
		ChatID:   c.ChatID,
		Username: c.Details.DisplayName,
		Content:  c.Details.DisplayName + " left the chat",
	}
}

// broadcast sends a message to all clients in a specific chat room.
// If the client's message buffer is full, the message will be dropped.
func (s *Server) broadcast(m *Message) {
	s.mutex.Lock()

	defer s.mutex.Unlock()

	chat, exists := s.Chats[m.ChatID]

	if !exists {
		log.Printf("Chat %s does not exist", m.ChatID)
		return
	}

	// Iterate over all clients in the chat
	// and attempt to send the message.
	for _, c := range chat.Clients {
		select {
		case c.Message <- m:
			log.Printf("Broadcast message to client %s: %s", c.Details.DisplayName, m.Content)
		default:
			log.Printf("Client %s message buffer full, dropping message", c.Details.DisplayName)
		}
	}
}
