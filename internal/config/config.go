package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
// Now includes JWTSecret for JWT authentication

type Config struct {
	DSN       string
	Port      string
	Mode      string
	JWTSecret string
}

// Load reads configuration from environment variables or sets defaults
func Load() Config {
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=demo port=5432 sslmode=disable"
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = ":8080"
	}

	mode := os.Getenv("APP_MODE")
	if mode == "" {
		mode = "debug"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your_jwt_secret"
	}

	return Config{
		DSN:       dsn,
		Port:      port,
		Mode:      mode,
		JWTSecret: jwtSecret,
	}
}
