package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	if len(os.Args) < 2 {
		log.Fatal("Please provide a migration direction: 'up' or 'down'.")
	}

	direction := os.Args[1]

	dsn := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create database driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	switch direction {
	case "up":
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to apply migration: %v", err)
		}
		log.Println("migration up success")
	case "down":
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("Failed to rollback migration: %v", err)
		}
		log.Println("migration down success")
	case "force":
		if len(os.Args) < 3 {
			log.Fatal("Please provide a migration version to force.")
		}

		versionStr := os.Args[2]
		version, err := strconv.ParseInt(versionStr, 10, 64)
		if err != nil {
			log.Fatalf("Invalid migration version: %v", err)
		}

		if err := m.Force(int(version)); err != nil {
			log.Fatalf("Failed to force migration %s: %v", versionStr, err)
		}
		log.Printf("Migration %s forced successfully", versionStr)
	default:
		log.Fatal("Invalid direction. Please use 'up' or 'down'.")
	}
}
