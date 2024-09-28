package main

import (
	"log"
	"time"

	"github.com/avila-r/social/e2e"
	"github.com/gofiber/fiber/v2"
)

var (
	api_url = e2e.Env.Get("API_URL")
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if err := app.Listen(api_url); err != nil {
		log.Fatal(err.Error())
	}

	// Wait for 5 seconds to assert that Fiber app is up before e2e test
	time.Sleep(5 * time.Second)
}
