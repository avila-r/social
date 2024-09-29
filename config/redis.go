package config

import (
	"log"

	application "github.com/avila-r/social"
	"github.com/joeshaw/envdecode"
)

// redis_properties holds the configuration properties for connecting to a Redis instance.
type redis_properties struct {
	// The URI for connecting to the Redis instance, required from the environment.
	URi string `env:"REDIS_URL,required"`
}

// RedisProperties is a singleton instance of redis_properties,
// initialized with the Redis configuration from environment variables.
var (
	RedisProperties = func() *redis_properties {
		// Asserts that app's environment is loaded
		application.Env.Load()

		var (
			p redis_properties // Create an instance of redis_pro
		)

		// Decode environment variables into the redis_properties instance.
		// If decoding fails, log a fatal error and terminate the application.
		if err := envdecode.StrictDecode(&p); err != nil {
			log.Fatalf("failed to decode redis's envs: %s", err)
		}

		// Return the populated redis_properties instance.
		return &p
	}()
)
