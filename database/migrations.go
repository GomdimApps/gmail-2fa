package database

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(databaseURL string, migrationsPath string) {
	m, err := migrate.New(
		"file://"+migrationsPath,
		databaseURL,
	)
	if err != nil {
		log.Fatalf("Error creating new migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	sourceErr, dbErr := m.Close()
	if sourceErr != nil {
		log.Printf("Warning: error closing migration source: %v", sourceErr)
	}
	if dbErr != nil {
		log.Printf("Warning: error closing migration database connection: %v", dbErr)
	}

	log.Println("Migrations applied successfully (if any were pending).")
}
