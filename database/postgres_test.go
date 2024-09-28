package database_test

import (
	"testing"

	application "github.com/avila-r/social"
	"github.com/avila-r/social/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func Test_PostgresConnection(t *testing.T) {
	// Asserts that application will load .env.test instead
	application.Env.ActivateTestEnvironment()

	// It's a good approach to start our Postgres'
	// instance dedicated to unit tests.
	//
	// Just run task `make test_start` to do that.
	psql := database.Postgres()

	// To generate SQL schema
	psql.AutoMigrate(&User{})

	u := User{
		Name: "test-name",
	}

	r := psql.Create(&u)

	if err := r.Error; err != nil {
		t.Error(err)
	}

	t.Logf("Created user's ID: %v", u.ID)
}
