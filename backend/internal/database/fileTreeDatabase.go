package database

import (
	"backend/internal/configuration"
	"errors"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO

type FileTreeDatabase struct {
	configuration      configuration.DbConfiguration
	DatabaseConnection *gorm.DB
}

func NewFileTreeDatabase() *FileTreeDatabase {
	return &FileTreeDatabase{}
}

func (fileTree *FileTreeDatabase) AddConfiguration(configuration configuration.DbConfiguration) *FileTreeDatabase {
	fileTree.configuration = configuration
	return fileTree
}

func (fileTreeDatabase *FileTreeDatabase) Build() (*FileTreeDatabase, error) {
	validationError := fileTreeDatabase.validateConfiguration()
	if validationError != nil {
		return nil, validationError
	}

	connectionString := fileTreeDatabase.configuration.GetConnectionString()
	databaseConnection, openDatabaseConnectionError := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if openDatabaseConnectionError != nil {
		log.Fatal(openDatabaseConnectionError)
	}

	fileTreeDatabase.DatabaseConnection = databaseConnection
	return fileTreeDatabase, nil
}

func (fileTreeDatabase *FileTreeDatabase) validateConfiguration() error {
	if fileTreeDatabase.configuration.Port <= 0 {
		return errors.New("configuration not complete: No 'Port' provided")
	}

	if fileTreeDatabase.configuration.Host == "" {
		return errors.New("configuration not complete: No 'Host' provided")
	}

	if fileTreeDatabase.configuration.Username == "" {
		return errors.New("configuration not complete: No 'Username' provided")
	}

	if fileTreeDatabase.configuration.Password == "" {
		return errors.New("configuration not complete: No 'Password' provided")
	}

	if fileTreeDatabase.configuration.DbName == "" {
		return errors.New("configuration not complete: No 'DbName' provided")
	}

	return nil
}
