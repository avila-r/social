package posts

import (
	"github.com/avila-r/xgo/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type (
	PostHandler struct {
		Service *PostService
	}
)

var (
	DefaultHandler = &PostHandler{
		Service: DefaultService,
	}
)

func NewHandler(s *PostService) *PostHandler {
	return &PostHandler{
		Service: s,
	}
}

func (h *PostHandler) Route(r fiber.Router) {
	// List all
	r.Get("/", func(c *fiber.Ctx) error {
		list, err := h.Service.ListPosts()

		if err != nil {
			return err
		}

		var (
			response []PostResponse
		)

		for _, post := range list {
			response = append(response, post.ToResponse())
		}

		return c.JSON(response)
	})

	// Create post
	r.Post("/", func(c *fiber.Ctx) error {
		var (
			request PostRequest
		)

		if err := c.BodyParser(&request); err != nil {
			return err
		}

		if err := validator.Validate(r); err != nil {
			response := validator.Map(err)

			return c.Status(500).JSON(response)
		}

		post := &Post{
			SenderID: request.Sender,
			Content:  request.Content,
		}

		if err := h.Service.CreatePost(post); err != nil {
			return err
		}

		return c.Status(201).JSON(post.ToResponse())
	})

	// Get by ID
	r.Get("/id/:id", func(c *fiber.Ctx) error {
		return nil
	})

	// Delete by ID
	r.Delete("/id/:id", func(c *fiber.Ctx) error {
		return nil
	})
}
