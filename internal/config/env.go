package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
	SecretKey   string
}

func LoadEnv() *Config {
	// CHECK ENV FILE
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file not found")
	}

	// DATABASE_URL
	dbURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok || dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// PORT
	port, ok := os.LookupEnv("PORT")
	if !ok || port == "" {
		log.Println("PORT is not set, using default value: 8080")
		port = "8080"
	}

	// JWT_SECRET
	secretKey, ok := os.LookupEnv("SECRET_KEY")
	if !ok || secretKey == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	return &Config{
		DatabaseURL: dbURL,
		Port:        port,
		SecretKey:   secretKey,
	}
}
