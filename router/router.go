package router

import (
	"github.com/avila-r/social/domain/auth"
	"github.com/avila-r/social/domain/users"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

var (
	Run = func(app *fiber.App) {
		// Public
		app.Get("/docs/*", swagger.HandlerDefault)

		// Check api connection
		app.Get("/verify", func(c *fiber.Ctx) error {
			return c.SendStatus(200)
		})

		auth.DefaultHandler.Route(
			app.Group("/auth"),
		)

		api := app.Group("/api")
		{
			api.Use(auth.Middleware)

			v1 := api.Group("/v1")
			{
				users.DefaultHandler.Route(
					v1.Group("/users"),
				)
			}
		}
	}
)
