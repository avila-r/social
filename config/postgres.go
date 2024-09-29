package config

import (
	"log"

	application "github.com/avila-r/social"
	"github.com/joeshaw/envdecode"
)

// postgres_properties holds the configuration properties for connecting to a PostgreSQL database.
type postgres_properties struct {
	DSN string `env:"POSTGRES_DSN,required"`
}

// PostgresProperties is a singleton instance of postgres_properties,
// initialized with the PostgreSQL database configuration from environment variables.
var (
	PostgresProperties = func() *postgres_properties {
		// Asserts that app's environment is loaded
		application.Env.Load()

		var (
			p postgres_properties // Create an instance of postgres_properties to hold the decoded configuration.
		)

		// Decode environment variables into the postgres_properties instance.
		// If decoding fails, log a fatal error and terminate the application.
		if err := envdecode.StrictDecode(&p); err != nil {
			log.Fatalf("failed to decode postgres' envs: %v", err.Error())
		}

		// Return the populated postgres_properties instance.
		return &p
	}()
)
