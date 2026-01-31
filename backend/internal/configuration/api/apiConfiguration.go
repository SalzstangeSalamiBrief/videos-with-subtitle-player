package api

import (
	"backend/internal/configuration"
	"fmt"
)

var defaultServerAddress = "localhost"
var defaultServerPort int64 = 3000

type ApiConfiguration struct {
	address     string
	port        int64
	allowedCors string
	rootPath    string
}

func NewApiConfiguration() (*ApiConfiguration, error) {
	rootPath, err := configuration.GetEnvironmentString("RootPath", true, nil)
	if err != nil {
		return nil, err
	}

	allowedCors, err := configuration.GetEnvironmentString("AllowedCors", true, nil)
	if err != nil {
		return nil, err
	}

	address, err := configuration.GetEnvironmentString("ServerAddress", false, &defaultServerAddress)
	if err != nil {
		return nil, err
	}

	port, err := configuration.GetEnvironmentInt("Port", false, &defaultServerPort)
	if err != nil {
		return nil, err
	}

	return &ApiConfiguration{
		address,
		port,
		allowedCors,
		rootPath,
	}, nil
}

func (configuration *ApiConfiguration) GetRootPath() string {
	return configuration.rootPath
}

func (configuration *ApiConfiguration) GetCors() string {
	return configuration.allowedCors
}

func (configuration *ApiConfiguration) GetServerAddress() string {
	return fmt.Sprintf("%v:%v", configuration.address, configuration.port)

}
