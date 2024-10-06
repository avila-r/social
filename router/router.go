package router

import (
	"github.com/avila-r/social/domain/auth"
	"github.com/avila-r/social/domain/chats"
	"github.com/avila-r/social/domain/posts"
	"github.com/avila-r/social/domain/users"

	"github.com/gofiber/fiber/v2"
)

// Run function initializes the application routes and middleware.
//
// Note that this method use only default handlers.
var (
	Run = func(app *fiber.App) {
		// Health check route to verify API connection
		app.Get("/verify", func(c *fiber.Ctx) error {
			return c.SendStatus(200)
		})

		// Routes for authentication-related endpoints
		auth.DefaultHandler.Route(
			app.Group("/auth"),
		)

		// Group for authenticated API routes
		api := app.Group("/api")
		{
			// Middleware to protect routes with authentication
			// api.Use(auth.Middleware)

			// Version 1 of the API
			v1 := api.Group("/v1")
			{
				// Routes for user-related operations
				users.DefaultHandler.Route(
					v1.Group("/users"),
				)

				// Routes for post-related operations
				posts.DefaultHandler.Route(
					v1.Group("/posts"),
				)

				// Group for WebSocket chat-related operations
				chat := v1.Group("/chats")
				{
					// Routes for chat-related operations
					chats.DefaultHandler.Route(
						chat,
					)
				}
			}
		}
	}
)
