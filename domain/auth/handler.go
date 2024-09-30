package auth

import (
	"github.com/avila-r/social/domain/users"
	"github.com/avila-r/xgo/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type (
	AuthHandler struct {
		Provider *XJwt
	}
)

var (
	DefaultHandler = &AuthHandler{
		Provider: DefaultService,
	}
)

func NewHandler(provider *XJwt) *AuthHandler {
	return &AuthHandler{
		Provider: provider,
	}
}

func (h *AuthHandler) Route(r fiber.Router) {
	r.Post("/login", func(c *fiber.Ctx) error {
		var (
			l LoginRequest
		)

		if err := c.BodyParser(&l); err != nil {
			return fiber.ErrBadRequest
		}

		response, err := h.Provider.Login(l)

		if err != nil {
			return err
		}

		return c.JSON(response)
	})

	// Create user
	r.Post("/register", func(c *fiber.Ctx) error {
		request := struct {
			DisplayName string `json:"display_name" validate:"required,min=3,max=50"`
			Email       string `json:"email"        validate:"required,email"`
			Password    string `json:"password"     validate:"required,min=8"`
		}{}

		if err := c.BodyParser(&request); err != nil {
			return fiber.ErrBadRequest
		}

		if err := validator.Validate(r); err != nil {
			response := validator.Map(err)

			return c.Status(500).JSON(response)
		}

		u := &users.User{
			DisplayName: request.DisplayName,
			Email:       request.Email,
			Password:    request.Password,
		}

		if err := h.Provider.UserService.CreateUser(u); err != nil {
			return err
		}

		return c.Status(201).JSON(u)
	})
}
