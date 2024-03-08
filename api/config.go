package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)
// Env returns TWILIO_ACCOUNT_SID, TWILIO_AUTHTOKEN and TWILIO_SERVICE_ID in order
func env(path string) (string, string, string) {
	fmt.Println(godotenv.Unmarshal(path))
	err := godotenv.Load(path)
	if err != nil {
		log.Fatalf("Error loading .env: %v", err)
	}
	return os.Getenv("TWILIO_ACCOUNT_SID"),
		os.Getenv("TWILIO_AUTHTOKEN"),
		os.Getenv("TWILIO_SERVICE_ID")

}
