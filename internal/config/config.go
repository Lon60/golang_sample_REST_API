package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DSN  string
	Port string
	Mode string
}

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
	return Config{
		DSN:  dsn,
		Port: port,
		Mode: mode,
	}
}
