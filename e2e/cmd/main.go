package main

import (
	"log"
	"time"

	application "github.com/avila-r/social"
	"github.com/gofiber/fiber/v2"
)

var (
	url = application.Env.Get("SERVER_URL")
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if err := app.Listen(url); err != nil {
		log.Fatal(err.Error())
	}

	// Wait for 5 seconds to assert that Fiber app is up before e2e test
	time.Sleep(5 * time.Second)
}
