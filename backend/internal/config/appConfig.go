package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	RootPath    string
	AllowedCors string
}

var AppConfiguration AppConfig

func InitializeConfiguration() {
	err := godotenv.Load()
	if err != nil {
		log.Default().Print("Could not load .env file; Use os arguments instead")
	}

	loadRootPath()
	loadAllowedCors()
}

func loadRootPath() {
	rp := os.Getenv("ROOT_PATH")
	if rp == "" {
		log.Fatal("Could not load environment variable ROOT_PATH")
	}

	AppConfiguration.RootPath = rp
}

func loadAllowedCors() {
	AppConfiguration.AllowedCors = os.Getenv("ALLOWED_CORS")
}
