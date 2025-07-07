package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Print("Error loading .env file. Using default values.")
	}
}

func GetEnv(key string, fallback string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		return fallback
	}
	
	return value
}