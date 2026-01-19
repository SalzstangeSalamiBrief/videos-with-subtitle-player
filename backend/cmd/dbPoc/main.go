package main

import (
	"backend/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	connectionString := "host=localhost port=5432 user=streamingServiceAccount password=tempPassword dbname=streaming-db sslmode=disable TimeZone=Europe/Berlin"
	databaseConnection, openDatabaseConnectionError := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if openDatabaseConnectionError != nil {
		log.Fatal(openDatabaseConnectionError)
	}

	db, dbError := databaseConnection.DB()
	defer db.Close()
	if dbError != nil {
		log.Fatal(openDatabaseConnectionError)
	}

	migrationError := databaseConnection.AutoMigrate(&models.FileTreeItem{})
	if migrationError != nil {
		log.Fatal(migrationError)
	}
}
