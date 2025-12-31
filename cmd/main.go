package main

import (
	"event-app/cmd/api"
	"event-app/internal/config"
)

func main() {
	cfg := config.LoadEnv()
	db := config.ConnectDatabse(cfg)

	server := api.NewServer(cfg, db)

	server.Run()
}
