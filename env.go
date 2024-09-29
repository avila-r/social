package application

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/avila-r/g"
	"github.com/avila-r/social/internal"
	"github.com/joho/godotenv"
)

type (
	environment struct {
		loaded    bool
		test      bool
		file_name string
	}
)

var (
	Env = environment{}
)

func (e *environment) Load() error {
	var (
		def  string
		file string
	)

	if e.test {
		def = ".env.test"
	} else {
		def = ".env"
	}

	file = g.If(e.file_name != "", e.file_name, def)

	path := filepath.Join(internal.RootPath, file)

	if err := godotenv.Load(path); err != nil {
		return err
	}

	e.loaded = true

	return nil
}

func (e *environment) Get(key string) string {
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

func (e *environment) SetFileName(n string) {
	if n != "" {
		e.file_name = n
	}
}

func (e *environment) ActivateTestEnvironment() {
	if f := !e.test; f {
		e.test = true
	}
}

func (e *environment) DeactivateTestEnvironment() {
	if t := e.test; t {
		e.test = false
	}
}
