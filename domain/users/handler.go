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
	r.Get("/", h.HandleHome)
}

func (h *UserHandler) HandleHome(c *fiber.Ctx) error {
	return c.JSON("Hello from user routes")
}