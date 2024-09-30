package users

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	DisplayName string
	Email       string
	Password    string
}
