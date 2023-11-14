package middlewares

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// DotEnvVariable -> get .env
func DotEnvVariable(key string) string {
	// Check if the application is running in a Docker environment
	if os.Getenv("DOCKER_ENV") == "true" {
		// In Docker, do not load .env file
		return os.Getenv(key)
	}

	// Load .env file if not in Docker environment
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	return os.Getenv(key)
}
