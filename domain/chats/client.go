package chats

import (
	"log"

	"github.com/avila-r/social/domain/users"
	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
)

type Client struct {
	Connection *websocket.Conn

	Details *users.User   `json:"client_details"`
	ChatID  uuid.UUID     `json:"chat_id"`
	Message chan *Message `json:"message"`
}

func (c *Client) Write() {
	defer func() {
		c.Connection.Close()

		log.Printf("Connection closed for client: %s", c.Details.DisplayName)
	}()

	for message := range c.Message {
		if err := c.Connection.WriteJSON(message); err != nil {
			log.Printf("Error sending message to client %s: %v", c.Details.DisplayName, err)
			return
		}
	}

	log.Printf("Message channel closed for client: %s", c.Details.DisplayName)
}

func (c *Client) Read(s *Server) {
	defer func() {
		s.Unregister <- c

		c.Connection.Close()

		log.Printf("Connection closed for client: %s", c.Details.DisplayName)
	}()

	for {
		_, body, err := c.Connection.ReadMessage()

		if err != nil {
			log.Printf("Error reading message from client %s: %v", c.Details.DisplayName, err)
			break
		}

		s.Broadcast <- &Message{
			ChatID:   c.ChatID,
			Username: c.Details.DisplayName,
			Content:  string(body),
		}
	}
}
