package main

import (
	"backend/pkg/models"
	"backend/pkg/services/fileTreeManager"
	"context"
	"database/sql"
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
	// TODO TRANSACTION WHILE MIGRATING
	db, dbError := databaseConnection.DB()
	defer db.Close()
	if dbError != nil {
		log.Fatal(openDatabaseConnectionError)
	}

	addFileTypeEnumError := getAndExecuteSqlFile(db, "1_AddFileTypeEnum.sql")
	if addFileTypeEnumError != nil {
		log.Fatal(addFileTypeEnumError)
	}

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

	seedTagsError := getAndExecuteSqlFile(db, "2_tags_seed.sql")
	if seedTagsError != nil {
		log.Fatal(seedTagsError)
	}

	fileTreeItemsFromDb, getFileTreeItemsFromDbError := gorm.G[models.FileTreeItem](databaseConnection).Find(ctx)
	if getFileTreeItemsFromDbError != nil {
		log.Fatal(getFileTreeItemsFromDbError)
	}

	log.Println(fileTreeItemsFromDb)
}

func getAndExecuteSqlFile(db *sql.DB, filename string) error {
	addFileTypeMigration, addFileTypeMigrationError := getMigrationFileContent(filename)
	if addFileTypeMigrationError != nil {
		return addFileTypeMigrationError
	}

	_, executeAddFileTypeMigrationError := db.Exec(addFileTypeMigration)
	if executeAddFileTypeMigrationError != nil {
		return executeAddFileTypeMigrationError
	}

	return nil
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
