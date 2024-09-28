package e2e_test

import (
	"testing"

	"github.com/avila-r/social/config"
	"github.com/avila-r/social/e2e"
)

func Test_Properties(t *testing.T) {
	// Assert that dotenv file is loaded before test
	e2e.Env.Load()

	_ = config.GetPostgresConfig()
}
