package util

import (
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnvVariable(key string) string {
	godotenv.Load(".env")
	return os.Getenv(key)
}
