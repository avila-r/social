package users

import (
	"github.com/avila-r/xgo/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type (
	UserHandler struct {
		Service *UserService
	}
)

var (
	DefaultHandler = &UserHandler{
		Service: DefaultService,
	}
)

func NewHandler(s *UserService) *UserHandler {
	return &UserHandler{
		Service: s,
	}
}

func (h *UserHandler) Route(r fiber.Router) {
	// Create user
	r.Post("/", func(c *fiber.Ctx) error {
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

		u := &User{
			DisplayName: request.DisplayName,
			Email:       request.Email,
			Password:    request.Password,
		}

		if err := h.Service.CreateUser(u); err != nil {
			return err
		}

		return c.Status(201).JSON(u)
	})

	// List users
	r.Get("/", func(c *fiber.Ctx) error {
		response, err := h.Service.ListUsers()

		if err != nil {
			return err
		}

		return c.JSON(response)
	})

	// Get by email
	r.Get("/email/:email", func(c *fiber.Ctx) error {
		email := c.Params("email")

		response, err := h.Service.FindByEmail(email)

		if err != nil {
			return err
		}

		return c.JSON(response)
	})
}
