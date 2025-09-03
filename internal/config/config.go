package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBName string
	DBHost string
	DBPort string
	DBUser string
	DBPass string
}

var Cfg = NewConfig()

func NewConfig() *Config {
	if err := godotenv.Load("internal/.env"); err != nil {
		log.Println("Error loading .env file, using defaults.")
	}
	return &Config{
		DBName: getEnv("DB_NAME", "url-shorteningdb"),
		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "5432"),
		DBUser: getEnv("DB_USER", "url-shortening"),
		DBPass: getEnv("DB_PASS", "123456"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
