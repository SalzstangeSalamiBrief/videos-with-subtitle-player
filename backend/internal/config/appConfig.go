package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const DEFAULT_HOST_ADDRESS = "localhost"
const DEFAULT_HOST_PORT = "3000"

type AppConfig struct {
	RootPath      string
	AllowedCors   string
	ServerAddress string
}

var AppConfiguration AppConfig

func InitializeConfiguration() AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Default().Print("Could not load .env file; Use os arguments instead")
	}

	loadServerAddress()
	loadRootPath()
	loadAllowedCors()
	return AppConfiguration
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

func loadServerAddress() {
	a := os.Getenv("HOST_ADDRESS")
	if a == "" {
		log.Default().Printf("Could not load environment variable HOST_ADDRESS. Use default value '%v'", DEFAULT_HOST_ADDRESS)
		a = DEFAULT_HOST_ADDRESS
	}

	p := os.Getenv("HOST_PORT")
	if p == "" {
		log.Default().Printf("Could not load environment variable HOST_PORT. Use default value '%v'", DEFAULT_HOST_PORT)
		p = DEFAULT_HOST_PORT
	}
	AppConfiguration.ServerAddress = fmt.Sprintf("%v:%v", a, p)
}
