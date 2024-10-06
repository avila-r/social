package posts

import (
	"gorm.io/gorm"

	"github.com/avila-r/social/database"
	"github.com/avila-r/social/domain/users"
)

type (
	PostService struct {
		Db *gorm.DB
		*users.UserService
	}
)

var (
	DefaultService = &PostService{
		Db:          database.Postgres,
		UserService: users.DefaultService,
	}
)

func NewService(db *gorm.DB) *PostService {
	return &PostService{
		Db:          db,
		UserService: users.DefaultService,
	}
}
