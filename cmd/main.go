package main

import (
	application "github.com/avila-r/social"
	"github.com/avila-r/xgo/pkg/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: api.ErrorHandler,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	url := application.Env.Get("SERVER_URL")

	app.Listen(url)
}
