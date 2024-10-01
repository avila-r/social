package users

import (
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
