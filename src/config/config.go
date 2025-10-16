package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	if os.Getenv("JWT_SECRET") == "" {
		log.Fatal("JWT_SECRET not set in .env")
	}
	if os.Getenv("DB_PATH") == "" {
		log.Fatal("DB_PATH not set in .env")
	}
}