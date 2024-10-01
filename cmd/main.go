package main

import (
	"github.com/avila-r/social/app"
	"github.com/avila-r/social/config"
)

func main() {
	// Load the server URL from the application configuration
	url := config.ApplicationConf.ServerURL

	// Start Fiber's app
	app.Instance.Listen(url)
}
