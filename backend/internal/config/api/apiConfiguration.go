package api

import (
	"errors"
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

func NewApiConfiguration() (*ApiConfiguration, error) {
	rootPath, err := loadRootPath()
	if err != nil {
		return nil, err
	}

	allowedCorsKey, err := loadAllowedCors()
	if err != nil {
		return nil, err
	}

	return &ApiConfiguration{
		rootPath:      rootPath,
		serverAddress: loadServerAddress(),
		allowedCors:   allowedCorsKey,
	}, nil
}

func loadRootPath() (string, error) {
	rp := os.Getenv(rootPathKey)
	if rp == "" {
		return "", errors.New(fmt.Sprintf("Could not load environment variable '%v'", rootPathKey))
	}

	_, checkPathExistsError := os.Stat(rp)
	if os.IsNotExist(checkPathExistsError) {
		return "", errors.New(fmt.Sprintf("The path '%v' does not exist", rp))
	}

	return rp, nil
}

func loadAllowedCors() (string, error) {
	allowedCors := os.Getenv(allowedCorsKey)
	if allowedCors == "" {
		return "", errors.New(fmt.Sprintf("Could not load environment variable '%v'", allowedCorsKey))

	}

	return allowedCors, nil
}

func loadServerAddress() string {
	address := os.Getenv(serverAddressKey)
	if address == "" {
		log.Default().Printf("Could not load environment variable '%v'. Use default value '%v'", serverAddressKey, defaultServerAddress)
		address = defaultServerAddress
	}

	port := os.Getenv(serverPortKey)
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
