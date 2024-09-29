package e2e_test

import (
	"testing"

	"github.com/avila-r/social/config"
)

func Test_Properties(t *testing.T) {
	_ = config.PostgresProperties
}
