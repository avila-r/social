package app

import (
	application "github.com/avila-r/social"
	"github.com/avila-r/social/router"
	"github.com/avila-r/xgo/pkg/api"

	"github.com/gofiber/fiber/v2"
)

var (
	Instance = func() *fiber.App {
		app := fiber.New(fiber.Config{
			// Custom error handling function
			//
			// Check https://github.com/avila-r/xgo/blob/main/pkg/api/errors.go to get more information
			ErrorHandler: api.ErrorHandler,
		})

		r := application.Env.Get("ROUTER")

		if r != "disable" {
			router.Run(app)
		}

		return app
	}()
)
