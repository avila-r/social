package application

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/avila-r/g"
	"github.com/joho/godotenv"
)

type (
	env struct {
		loaded bool
	}
)

var (
	Env = env{}
)

func (e *env) Load() error {
	if e.loaded {
		return nil
	}

	def := filepath.Join(RootPath, ".env")

	// Loads default .env
	if err := godotenv.Load(def); err != nil {
		return err
	}

	profile := os.Getenv("PROFILE")

	custom := ".env" + g.If(profile == "", profile, "."+profile)

	if profile == "e2e" {
		custom = filepath.Join(E2ePath, ".env")
	}

	path := filepath.Join(RootPath, custom)

	if err := godotenv.Load(path); err != nil {
		return err
	}

	e.loaded = true

	return nil
}

func (e *env) Get(key string) string {
	var (
		once sync.Once
	)

	if !e.loaded {
		if err := e.Load(); err != nil {
			once.Do(func() {
				log.Print("warning: there is no .env file to load.")
			})
		}
	}

	return os.Getenv(key)
}
