package auth_test

import (
	"testing"

	"github.com/avila-r/social/domain/auth"
	"github.com/avila-r/social/domain/users"
)

var (
	user = &users.User{
		DisplayName: "Alice Test",
		Email:       "alice.test@example.com",
		Password:    "password123",
	}

	login = auth.LoginRequest{
		Email:    "alice.test@example.com",
		Password: "password123",
	}
)

func Test_Login(t *testing.T) {
	// Cleanup: Delete the test user after the test
	t.Cleanup(func() {
		users.DefaultService.Db.Where("id = ?", user.ID).Delete(user)
	})

	if err := users.DefaultService.CreateUser(user); err != nil {
		t.Errorf(err.Error())
	}

	if _, err := auth.DefaultService.Login(login); err != nil {
		t.Errorf(err.Error())
	}
}
