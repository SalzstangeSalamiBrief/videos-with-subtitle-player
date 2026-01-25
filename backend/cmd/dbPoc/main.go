package main

import (
	"backend/pkg/models"
	"backend/pkg/services/fileTreeManager"
	"context"
	"log"
	"os"
	"path"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	firstMigrationResult, executeFirstMigrationError := db.Exec(firstMigration)
	if executeFirstMigrationError != nil {
		log.Fatal(executeFirstMigrationError)
	}

	log.Println(firstMigrationResult)

	migrationError := databaseConnection.AutoMigrate(&models.FileTreeItem{}, &models.Tag{})
	if migrationError != nil {
		log.Fatal(migrationError)
	}

	initializedFileTreeManager := fileTreeManager.NewFileTreeManager("E:\\projects\\videos-with-subtitle-player\\test-content").InitializeTree()
	fileTree := initializedFileTreeManager.GetTree()
	_, clearDbError := db.Exec("DELETE FROM file_tree_items")
	if clearDbError != nil {
		log.Fatal(clearDbError)
	}

	result := gorm.WithResult()
	ctx := context.Background()
	createInBatchError := gorm.G[models.FileTreeItem](databaseConnection, result).CreateInBatches(ctx, &fileTree, len(fileTree))
	if createInBatchError != nil {
		log.Fatal(createInBatchError)
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
