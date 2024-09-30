package users

import (
	"github.com/avila-r/social/database"

	"gorm.io/gorm"
)

type (
	UserService struct {
		Db *gorm.DB
	}
)

var (
	DefaultService = func() *UserService {
		return &UserService{
			Db: database.Postgres,
		}
	}()
)

func NewService(db *gorm.DB) *UserService {
	return &UserService{
		Db: db,
	}
}
