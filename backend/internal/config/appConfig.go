package config

import "os"

type AppConfig struct {
	RootPath string
}

var AppConfiguration AppConfig

func InitializeConfiguration() {
	AppConfiguration.RootPath = os.Getenv("ROOT_PATH")
}
