package users_test

import (
	"testing"

	u "github.com/avila-r/social/domain/users"
)

var (
	users = []*u.User{
		{
			DisplayName: "Alice Test",
			Email:       "alice.test@example.com",
			Password:    "password123",
		},
		{
			DisplayName: "Bob Test",
			Email:       "bob.test@example.com",
			Password:    "password456",
		},
		{
			DisplayName: "Charlie Test",
			Email:       "charlie.test@example.com",
			Password:    "password789",
		},
	}
)

func Test_CreateUser(t *testing.T) {
	service := u.DefaultService

	service.Db.AutoMigrate(&u.User{})

	user := &u.User{
		DisplayName: "Test display name",
		Email:       "avila.dev@outlook.com",
		Password:    "test-password",
	}

	// Cleanup: Delete the test user after the test
	t.Cleanup(func() {
		service.Db.Where("id = ?", user.ID).Delete(user)
	})

	if err := service.CreateUser(user); err != nil {
		t.Errorf("error while trying to create user - %v", err)
	}

	t.Run("when email is not unique", func(t *testing.T) {
		if err := service.CreateUser(user); err == nil {
			t.Errorf("no error was returned when creating equal user - %v", err)
		}
	})
}

func Test_FindByEmail(t *testing.T) {
	service := u.DefaultService

	// Migrate the User table
	service.Db.AutoMigrate(&u.User{})

	// Create a new user for testing
	user := &u.User{
		DisplayName: "Test display name",
		Email:       "avila.dev@outlook.com",
		Password:    "test-password",
	}

	// Cleanup: Delete the test user after the test
	t.Cleanup(func() {
		service.Db.Where("id = ?", user.ID).Delete(user)
	})

	// Create the user
	if err := service.CreateUser(user); err != nil {
		t.Errorf("error while trying to create user - %v", err)
	}

	// Attempt to find the created user by email
	if r, err := service.FindByEmail(user.Email); r == nil || err != nil {
		t.Errorf("error while trying to find user - %v", err)
	}
}

func Test_ListAll(t *testing.T) {
	service := u.DefaultService

	// Migrate the User table
	service.Db.AutoMigrate(&u.User{})

	// Cleanup: Delete the test user after the test
	t.Cleanup(func() {
		for _, user := range users {
			service.Db.Where("id = ?", user.ID).Delete(user)
		}
	})

	for _, user := range users {
		// Create the user
		if err := service.CreateUser(user); err != nil {
			t.Errorf("error while trying to create user - %v", err)
		}
	}

	_, err := service.ListUsers()

	if err != nil {
		t.Errorf("error while trying to list users - %v", err)
	}
}
