package posts

import (
	"time"

	"github.com/avila-r/social/domain/users"

	"github.com/google/uuid"
)

type Post struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	SenderID uuid.UUID
	Sender   *users.User `gorm:"foreinKey:SenderID;"`

	Content string
}

type PostRequest struct {
	Sender  uuid.UUID `json:"sender"  validate:"required, uuid"`
	Content string    `json:"content" validate:"required"`
}

type PostResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Sender    uuid.UUID `json:"sender"`
	Content   string    `json:"content"`
}

func (p *Post) ToResponse() PostResponse {
	return PostResponse{
		ID:        p.ID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Sender:    p.SenderID,
		Content:   p.Content,
	}
}
