package loadenv

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Errored while loading the .env file")
	}

	apikey := os.Getenv("COINCAP-API-KEY")
	return apikey
}
