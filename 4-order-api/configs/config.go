package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *Config {
	// In a real application, you would load this from a file or environment variables
	err := godotenv.Load() // Assuming you have a .env file for configuration

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"), // Load DSN from environment variable
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"), // Load Auth Secret from environment variable
		},
	}
}
