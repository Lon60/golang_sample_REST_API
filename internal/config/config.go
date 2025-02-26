package config

import "os"

type Config struct {
	DSN  string
	Port string
}

func Load() Config {
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=demo port=5432 sslmode=disable"
	}
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = ":8080"
	}
	return Config{
		DSN:  dsn,
		Port: port,
	}
}
