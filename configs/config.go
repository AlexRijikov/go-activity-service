package config

import (
	"log"
	"os"
)

type Config struct {
	AppPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func Load() *Config {
	cfg := &Config{
		AppPort: getEnv("APP_PORT", "8080"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "activity"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Printf("ENV %s not set, using default: %s", key, defaultValue)
		return defaultValue
	}
	return val
}
