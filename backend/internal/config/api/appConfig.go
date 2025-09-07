package api

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const DEFAULT_HOST_ADDRESS = "localhost"
const DEFAULT_HOST_PORT = "3000"

// TODO REFACTOR: ONE CONFIG READER FOR API AND ONE FOR THE WEBP CMD
type configuration struct {
	RootPath      string
	AllowedCors   string
	ServerAddress string
}

var appConfiguration configuration

func InitializeConfiguration() configuration {
	err := godotenv.Load()
	if err != nil {
		log.Default().Print("Could not load .env file; Use os arguments instead")
	}

	loadServerAddress()
	loadRootPath()
	loadAllowedCors()
	return appConfiguration
}

func loadRootPath() {
	rp := os.Getenv("ROOT_PATH")
	if rp == "" {
		log.Fatal("Could not load environment variable ROOT_PATH")
	}

	appConfiguration.RootPath = rp
}

func loadAllowedCors() {
	appConfiguration.AllowedCors = os.Getenv("ALLOWED_CORS")
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
	appConfiguration.ServerAddress = fmt.Sprintf("%v:%v", a, p)
}
