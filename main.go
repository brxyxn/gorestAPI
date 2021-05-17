package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	a := App{}

	a.Initialize(
		dotEnvGet("BASICAPI_DB_USERNAME"),
		dotEnvGet("BASICAPI_DB_PASSWORD"),
		dotEnvGet("BASICAPI_DB_NAME"))

	a.Run(":3005")
}

// Get environment key from .env file
func dotEnvGet(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
