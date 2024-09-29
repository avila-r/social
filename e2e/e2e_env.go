package e2e

import (
	"errors"
	"os"
	"path/filepath"

	application "github.com/avila-r/social"
	"github.com/joho/godotenv"
)

type (
	environment struct {
		loaded bool
	}
)

var (
	Env = environment{}
)

func (e *environment) Load() error {
	path := filepath.Join(application.E2ePath, ".env")

	if err := godotenv.Load(path); err != nil {
		return err
	}

	e.loaded = true

	return nil
}

func (e *environment) Get(key string) string {
	if !e.loaded {
		if err := e.Load(); err != nil {
			panic(errors.New("there is no .env file to load"))
		}
	}

	return os.Getenv(key)
}
