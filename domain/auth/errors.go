package auth

import (
	"github.com/avila-r/xgo/pkg/api"
)

var (
	ErrUnexpectedSigningMethod = &api.Error{
		Code:    400,
		Message: "unexpected jwt signing method",
	}

	ErrUnauthorized = &api.Error{
		Code:    401,
		Message: "invalid or missing jwt",
	}

	ErrIncorrectLogin = &api.Error{
		Code:    401,
		Message: "invalid email or password",
	}
)
