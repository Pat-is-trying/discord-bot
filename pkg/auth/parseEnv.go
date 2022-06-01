package auth

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func FetchEnv(key string) (string) {
	err := godotenv.Load("../local.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}