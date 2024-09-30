package users_test

import (
	"testing"

	u "github.com/avila-r/social/domain/users"
)

var (
	users = []u.User{
		{
			DisplayName: "Test display name",
			Email:       "avila.dev@outlook.com",
			Password:    "test-password",
		},
		{
			DisplayName: "Test display name",
			Email:       "avila.dev@outlook.com",
			Password:    "test-password",
		},
		{
			DisplayName: "Test display name",
			Email:       "avila.dev@outlook.com",
			Password:    "test-password",
		},
	}
)

func Test_CreateUser(t *testing.T) {
	service := u.DefaultService

	service.Db.AutoMigrate(&u.User{})

	u := u.User{
		DisplayName: "Test display name",
		Email:       "avila.dev@outlook.com",
		Password:    "test-password",
	}

	_, err := service.CreateUser(u)

	if err != nil {
		t.Errorf("error while trying to create user - %v", err)
	}

	// t.Run("when email is not unique", func(t *testing.T) {
	// 	_, _ = service.CreateUser(u)

	// 	if _, err := service.CreateUser(u); err == nil {
	// 		t.Errorf("no error was returned when creating equal user - %v", err)
	// 	}
	// })
}

func Test_ListAll(t *testing.T) {
	service := u.DefaultService

	service.Db.AutoMigrate(&u.User{})

	for _, user := range users {
		_, err := service.CreateUser(user)

		if err != nil {
			t.Errorf("error while trying to create user - %v", err)
		}
	}

	result, err := service.ListUsers()

	if err != nil {
		t.Errorf("error while trying to list users - %v", err)
	}

	if len(result) == 0 {
		t.Errorf("no users to list")
	}
}
