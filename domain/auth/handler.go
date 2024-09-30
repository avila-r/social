package auth

import "github.com/gofiber/fiber/v2"

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
			l Login
		)

		if err := c.BodyParser(&l); err != nil {
			return fiber.ErrBadRequest
		}

		token, err := h.Provider.Login(l)

		if err != nil {
			return err
		}

		return c.JSON(fiber.Map{
			"token": token,
		})
	})
}
