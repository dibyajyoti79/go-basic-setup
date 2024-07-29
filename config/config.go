package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() error {
	return godotenv.Load()
}

func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}
