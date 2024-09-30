package auth

import (
	application "github.com/avila-r/social"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v5"
)

var (
	Middleware = jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Respond with a 401 Unauthorized status
			return ErrUnauthorized
		},

		SigningKey: jwtware.SigningKey{
			JWTAlg: "HS256",
			Key:    Secret,
		},
	})

	Secret = []byte(application.Env.Get("JWT_SECRET"))

	SigningMethod = jwt.SigningMethodHS256
)
