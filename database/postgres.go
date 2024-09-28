package database

import (
	"log"

	"github.com/avila-r/social/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Postgres = func() *gorm.DB {
		psql_props := config.GetPostgresConfig()

		dsn := psql_props.DSN

		db, err := gorm.Open(postgres.Open(dsn))

		if err != nil {
			log.Fatalf("failed to connect to database - %v", err)
		}

		AutoMigrate(db)

		return db
	}
)

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate()

	if err != nil {
		log.Fatalf(err.Error())
	}
}
