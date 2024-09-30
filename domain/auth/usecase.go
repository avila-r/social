package auth

import (
	"time"

	"github.com/avila-r/social/domain/users"
	"github.com/avila-r/xgo/pkg/crypt"
	"github.com/golang-jwt/jwt/v5"
)

type (
	XJwt struct {
		users.UserService
	}
)

var (
	DefaultService = &XJwt{
		UserService: *users.DefaultService,
	}
)

func (s *XJwt) Login(request Login) (string, error) {
	user, err := s.UserService.FindByEmail(request.Email)

	if err != nil {
		return "", ErrIncorrectLogin
	}

	if valid := crypt.PasswordMatchesHash(request.Password, user.Password); !valid {
		return "", ErrIncorrectLogin
	}

	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	return s.Generate(claims)
}

func (s *XJwt) Generate(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(SigningMethod, claims)

	return token.SignedString(Secret)
}
