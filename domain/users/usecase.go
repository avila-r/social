package users

import (
	"github.com/avila-r/xgo/pkg/crypt"
)

func (s *UserService) CanCreate(u *User) (bool, error) {
	// Check if the email already exists
	if _, err := s.FindByEmail(u.Email); err == nil {
		// Err is nil. This means that email already
		// exists, so it's not unique
		return false, ErrEmailNotUnique
	}

	// Everything is ok
	return true, nil
}

func (s *UserService) ListUsers() ([]User, error) {
	var (
		users []User
	)

	result := s.Db.Find(&users)

	return users, result.Error
}

func (s *UserService) CreateUser(user *User) error {
	hash, err := crypt.EncryptPassword(user.Password)

	if err != nil {
		return err
	}

	if valid, err := s.CanCreate(user); !valid {
		return err
	}

	user.Password = hash

	result := s.Db.Create(&user)

	return result.Error
}

// FindByEmail finds a user by their email address
func (s *UserService) FindByEmail(email string) (*User, error) {
	var (
		user User
	)

	// Search for user by email in the database
	result := s.Db.Where("email = ?", email).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
