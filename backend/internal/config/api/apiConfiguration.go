package api

import (
	"fmt"
	"log"
	"os"
)

const defaultServerAddress = "localhost"
const defaultServerPort = "3000"

const rootPathKey string = "RootPath"
const allowedCorsKey string = "allowedCors"
const serverAddressKey string = "serverAddress"
const serverPortKey string = "serverPort"

type ApiConfiguration struct {
	rootPath      string
	allowedCors   string
	serverAddress string
}

func NewApiConfiguration() *ApiConfiguration {
	return &ApiConfiguration{
		rootPath:      loadRootPath(),
		serverAddress: loadServerAddress(),
		allowedCors:   loadAllowedCors(),
	}
}

func loadRootPath() string {
	rp := os.Getenv(rootPathKey)
	if rp == "" {
		log.Fatalf("Could not load environment variable '%v'", rootPathKey)
	}

	return rp
}

func loadAllowedCors() string {
	allowedCors := os.Getenv(allowedCorsKey)
	if allowedCors == "" {
		log.Fatalf("Could not load environment variable '%v'", allowedCorsKey)

	}

	return allowedCors
}

func loadServerAddress() string {
	address := os.Getenv("HOST_ADDRESS")
	if address == "" {
		log.Default().Printf("Could not load environment variable '%v'. Use default value '%v'", serverAddressKey, defaultServerAddress)
		address = defaultServerAddress
	}

	port := os.Getenv("HOST_PORT")
	if port == "" {
		log.Default().Printf("Could not load environment variable '%v. Use default value '%v'", serverPortKey, defaultServerPort)
		port = defaultServerPort
	}

	return fmt.Sprintf("%v:%v", address, port)
}

func (configuration *ApiConfiguration) GetRootPath() string {
	return configuration.rootPath
}

func (configuration *ApiConfiguration) GetCors() string {
	return configuration.allowedCors
}

func (configuration *ApiConfiguration) GetServerAddress() string {
	return configuration.serverAddress
}
