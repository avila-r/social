package auth

import (
	"github.com/avila-r/social/domain/users"
)

type (
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		Token       string      `json:"access_token"`
		UserDetails *users.User `json:"user_details"`
	}
)
