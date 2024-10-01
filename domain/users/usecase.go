package users

import (
	"github.com/avila-r/xgo/pkg/crypt"
	"github.com/google/uuid"
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

	result := s.Db.Create(user)

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

func (s *UserService) FindByID(id uuid.UUID) (*User, error) {
	var (
		user User // Holds the user post
	)

	// Query the database for the first user with the given ID
	result := s.Db.Where("id = ?", id).First(&user)

	// Return the found user and any error
	// that occurred during the query
	return &user, result.Error
}

func (s *UserService) DeleteByID(id uuid.UUID) error {
	// Find the user by ID
	user, err := s.FindByID(id)

	if err != nil {
		// If an error occurred
		// (e.g., user not found),
		// return the error
		return err
	}

	// Delete the found user from the database
	result := s.Db.Delete(&user)

	// Return any error that occurred during the deletion
	return result.Error
}

func (s *UserService) ExistsById(id uuid.UUID) bool {
	var (
		user *User
	)

	// Search for user by ID in the database
	result := s.Db.Where("id = ?", id).First(&user)

	if result.Error != nil || user == nil {
		return false
	}

	return true
}
