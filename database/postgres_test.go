package database_test

import (
	"testing"

	application "github.com/avila-r/social"
	"github.com/avila-r/social/database"
	"gorm.io/gorm"
)

// User represents a model for the user entity with a name.
type User struct {
	gorm.Model
	Name string
}

// Test_PostgresConnection tests the connection to the PostgreSQL database and the ability to create a user.
func Test_PostgresConnection(t *testing.T) {
	// Asserts that application will load .env.test instead
	application.Env.ActivateTestEnvironment()

	// Retrieve the PostgreSQL database instance
	//
	// It's a good approach to start our PostgreSQL instance dedicated to unit tests.
	// Just run task `make test_start` to set up the test environment.
	psql := database.Postgres

	// Automatically migrate the User model to generate the SQL schema in the database.
	psql.AutoMigrate(&User{})

	// Create a new User instance for testing.
	u := User{
		Name: "test-name",
	}

	// Attempt to create the user in the database.
	r := psql.Create(&u)

	// Check for errors during the user creation process.
	if err := r.Error; err != nil {
		t.Error(err)
	}

	// Log the ID of the created user.
	t.Logf("Created user's ID: %v", u.ID)
}
