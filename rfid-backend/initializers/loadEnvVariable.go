package initializers

import (
	"github.com/joho/godotenv"
)

func LoadEnvVariable() {
	godotenv.Load()
}