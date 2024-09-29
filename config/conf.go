package config

import (
	"log"

	application "github.com/avila-r/social"
	"github.com/joeshaw/envdecode"
)

// application_conf holds the configuration properties for the application.
type application_conf struct {
	ServerURL string `env:"SERVER_URL,required"`
}

// ApplicationConf is a singleton instance of application_conf,
// initialized with the application's environment configuration.
var (
	ApplicationConf = func() *application_conf {
		// Asserts that app's environment is loaded
		application.Env.Load()

		var (
			a application_conf // Create an instance of application_conf to hold the decoded configuration.
		)

		// Decode environment variables into the application_conf instance.
		// If decoding fails, log a fatal error and terminate the application.
		if err := envdecode.StrictDecode(&a); err != nil {
			log.Fatalf("failed to decode app conf's envs - %v", err.Error())
		}

		// Return the populated application_conf instance.
		return &a
	}()
)
