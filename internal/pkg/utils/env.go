package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv set env vars from .env file for development mode
func LoadEnv() {
	env := os.Getenv("ENV")

	if env == "development" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}

		err = godotenv.Load(wd + "/.env")

		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}
}
