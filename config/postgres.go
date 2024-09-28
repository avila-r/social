package config

import (
	"log"

	application "github.com/avila-r/social"
	"github.com/joeshaw/envdecode"
)

type PostgresProps struct {
	DSN string `env:"POSTGRES_DSN,required"`
}

func GetPostgresConfig() *PostgresProps {
	// Asserts that app's environment is loaded
	application.Env.Load()

	var (
		c PostgresProps
	)

	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("failed to decode envs: %s", err)
	}

	return &c
}
