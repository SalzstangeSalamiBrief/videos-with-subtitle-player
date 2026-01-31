package api

import (
	"backend/internal/configuration"
	"fmt"
)

var defaultPort int64 = 5432

type DbConfiguration struct {
	host     string
	port     int64
	username string
	password string
	dbname   string
}

func NewDbConfiguration() (*DbConfiguration, error) {
	host, err := configuration.GetEnvironmentString("Host", true, nil)
	if err != nil {
		return nil, err
	}

	port, err := configuration.GetEnvironmentInt("Port", false, &defaultPort)
	if err != nil {
		return nil, err
	}

	username, err := configuration.GetEnvironmentString("Username", true, nil)
	if err != nil {
		return nil, err
	}

	password, err := configuration.GetEnvironmentString("Password", true, nil)
	if err != nil {
		return nil, err
	}

	dbname, err := configuration.GetEnvironmentString("DbName", true, nil)
	if err != nil {
		return nil, err
	}

	return &DbConfiguration{
		host,
		port,
		username,
		password,
		dbname,
	}, nil
}

func (configuration *DbConfiguration) GetHost() string {
	return configuration.host
}

func (configuration *DbConfiguration) GetPort() int64 {
	return configuration.port
}

func (configuration *DbConfiguration) GetUsername() string {
	return configuration.username
}

func (configuration *DbConfiguration) GetPassword() string {
	return configuration.password
}

func (configuration *DbConfiguration) GetDbName() string {
	return configuration.dbname
}

func (configuration *DbConfiguration) GetConnectionString() string {
	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable TimeZone=Europe/Berlin", configuration.host, configuration.port, configuration.username, configuration.password, configuration.dbname)
}
