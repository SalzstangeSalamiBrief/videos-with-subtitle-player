package config

import "os"

type AppConfig struct {
	RootPath string
}

var AppConfiguration AppConfig
// TODO build a custom solution that reads an env file
// TODO the custom solution generates a struct for the content and type safety
func InitializeConfiguration() {
	AppConfiguration.RootPath = os.Getenv("ROOT_PATH")
}
