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

func (s *XJwt) Login(request LoginRequest) (*LoginResponse, error) {
	user, err := s.UserService.FindByEmail(request.Email)

	if err != nil {
		return nil, ErrIncorrectLogin
	}

	if valid := crypt.PasswordMatchesHash(request.Password, user.Password); !valid {
		return nil, ErrIncorrectLogin
	}

	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token, err := s.Generate(claims)

	if err != nil {
		return nil, err
	}

	response := &LoginResponse{
		Token:       token,
		UserDetails: user,
	}

	return response, nil
}

func (s *XJwt) Generate(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(SigningMethod, claims)

	return token.SignedString(Secret)
}
