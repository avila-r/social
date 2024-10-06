package chats

import "github.com/google/uuid"

type Chat struct {
	ID      uuid.UUID             `json:"id"`
	Clients map[uuid.UUID]*Client `json:"clients"`
}

type Message struct {
	ChatID   uuid.UUID `json:"chat_id"`
	Username string    `json:"username"`
	Content  string    `json:"content"`
}
