package configuration

import (
	"backend/internal/configuration/utilities"
	"fmt"
)

var defaultPort int64 = 5432

type DbConfiguration struct {
	Port     int64
	Host     string
	Username string
	Password string
	DbName   string
}

func NewDbConfiguration() (*DbConfiguration, error) {
	host, err := utilities.GetEnvironmentString("DbHost", true, nil)
	if err != nil {
		return nil, err
	}

	port, err := utilities.GetEnvironmentInt("DbPort", false, &defaultPort)
	if err != nil {
		return nil, err
	}

	username, err := utilities.GetEnvironmentString("DbUsername", true, nil)
	if err != nil {
		return nil, err
	}

	password, err := utilities.GetEnvironmentString("DbPassword", true, nil)
	if err != nil {
		return nil, err
	}

	dbname, err := utilities.GetEnvironmentString("DbName", true, nil)
	if err != nil {
		return nil, err
	}

	return &DbConfiguration{
		port,
		host,
		username,
		password,
		dbname,
	}, nil
}

func (configuration *DbConfiguration) GetHost() string {
	return configuration.Host
}

func (configuration *DbConfiguration) GetPort() int64 {
	return configuration.Port
}

func (configuration *DbConfiguration) GetUsername() string {
	return configuration.Username
}

func (configuration *DbConfiguration) GetPassword() string {
	return configuration.Password
}

func (configuration *DbConfiguration) GetDbName() string {
	return configuration.DbName
}

func (configuration *DbConfiguration) GetConnectionString() string {
	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable TimeZone=Europe/Berlin", configuration.Host, configuration.Port, configuration.Username, configuration.Password, configuration.DbName)
}
