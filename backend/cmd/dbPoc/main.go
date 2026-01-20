package main

import (
	"backend/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"path"
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

	firstMigration, getFirstMigrationError := getMigrationFileContent("1_AddFileTypeEnum.sql")
	if getFirstMigrationError != nil {
		log.Fatal(getFirstMigrationError)
	}

	result, executeFirstMigrationError := db.Exec(firstMigration)
	if executeFirstMigrationError != nil {
		log.Fatal(executeFirstMigrationError)
	}

	log.Println(result)

	migrationError := databaseConnection.AutoMigrate(&models.FileTreeItem{})
	if migrationError != nil {
		log.Fatal(migrationError)
	}
}

func getMigrationFileContent(filename string) (string, error) {
	var relativeFilePathToMigrationFolder = path.Join("./migrations/", filename)
	contentAsByteArray, readFileError := os.ReadFile(relativeFilePathToMigrationFolder)
	if readFileError != nil {
		return "", readFileError
	}

	var stringifiedContent = string(contentAsByteArray)
	return stringifiedContent, nil

}
