package chats

import (
	"log"
	"sync"

	"github.com/avila-r/social/domain/users"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	"github.com/google/uuid"
)

type (
	ChatHandler struct {
		Server *Server
		*users.UserService
		mutex sync.Mutex
	}
)

var (
	DefaultHandler = func() *ChatHandler {
		var (
			s *Server
		)

		var once sync.Once

		once.Do(func() {
			s = &Server{
				Chats:      make(map[uuid.UUID]*Chat),
				Register:   make(chan *Client),
				Unregister: make(chan *Client),

				// Buffered channel with
				// capacity of 5 messages
				Broadcast: make(chan *Message, 5),
			}
		})

		return &ChatHandler{
			Server: s,
		}
	}()
)

func NewHandler(s *Server) *ChatHandler {
	return &ChatHandler{
		Server: s,
	}
}

func (h *ChatHandler) Route(r fiber.Router) {
	// Create chat
	r.Post("/", func(c *fiber.Ctx) error {
		id, err := uuid.FromBytes([]byte(c.Params("id")))

		if err != nil {
			return err
		}

		h.mutex.Lock()

		defer h.mutex.Unlock()

		if _, exists := h.Server.Chats[id]; exists {
			return ErrChatAlreadyExists
		}

		created := &Chat{
			ID:      id,
			Clients: make(map[uuid.UUID]*Client),
		}

		h.Server.Chats[id] = created

		return c.Status(201).JSON(created)
	})

	// Join chat
	r.Get("/:id/connect", func(c *fiber.Ctx) error {
		if !websocket.IsWebSocketUpgrade(c) {
			return ErrWebsocketRequired
		}

		return websocket.New(func(c *websocket.Conn) {
			var (
				chat, _ = uuid.FromBytes(
					[]byte(c.Params("chat_id")),
				)

				user_id, _ = uuid.FromBytes(
					[]byte(c.Query("user_id")),
				)
			)

			user, err := h.UserService.FindByID(user_id)

			if err != nil {
				return
			}

			client := &Client{
				Connection: c,
				Details:    user,
				ChatID:     chat,
				Message:    make(chan *Message, 10),
			}

			h.Server.Register <- client

			h.Server.Broadcast <- &Message{
				ChatID:   chat,
				Username: user.DisplayName,
				Content:  user.DisplayName + " joined the chat",
			}

			log.Printf("User %s joined chat %s", user.DisplayName, chat)

			go client.Write()

			client.Read(h.Server)
		})(c)
	})
}
