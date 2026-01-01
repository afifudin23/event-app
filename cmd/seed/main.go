package main

import (
	"event-app/cmd/seed/seeders"
	"event-app/internal/config"
	"log"
)

func main() {
	log.Println("--- Starting Seeder Program ---")

	cfg := config.LoadEnv()
	db := config.ConnectDatabse(cfg)

	seeders.RunAllSeeder(db)

	log.Println("--- Seeder Program Finished Successfully ---")
}
