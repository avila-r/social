package database

import (
	"log"

	"github.com/avila-r/social/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// Singleton instance of the GORM database connection.
	Postgres = func() *gorm.DB {
		// Retrieve PostgreSQL properties from the configuration
		psql_props := config.PostgresProperties

		// Get the Data Source Name (DSN) for connecting to the database
		dsn := psql_props.DSN

		// Open a connection to the PostgreSQL database using GORM
		db, err := gorm.Open(postgres.Open(dsn))

		if err != nil {
			log.Fatalf("failed to connect to database - %v", err)
		}

		// Automatically migrate the database schema
		AutoMigrate(db)

		return db // Return the database instance
	}()
)

// AutoMigrate handles the automatic migration of the database schema.
func AutoMigrate(db *gorm.DB) {
	// Perform the auto migration for all registered models
	err := db.AutoMigrate()

	if err != nil {
		log.Fatalf(err.Error())
	}
}
