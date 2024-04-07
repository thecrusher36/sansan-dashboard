package libs

import (
	"os"

	"github.com/joho/godotenv"
)

// GetEnv return environment variable from .env file. return fallback string if not found
func GetEnv(key, fallback string) string {
	godotenv.Load(".env")
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
