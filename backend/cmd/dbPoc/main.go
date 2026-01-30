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
	fileTreeFromDisk := initializedFileTreeManager.GetTree()
	ctx := context.Background()

	seedTagsError := getAndExecuteSqlFile(db, "2_tags_seed.sql")
	if seedTagsError != nil {
		log.Fatal(seedTagsError)
	}

	fileTreeItemsFromDb, getFileTreeItemsFromDbError := gorm.G[models.FileTreeItem](databaseConnection).Find(ctx)
	if getFileTreeItemsFromDbError != nil {
		log.Fatal(getFileTreeItemsFromDbError)
	}

	syncError := syncFiles(databaseConnection, fileTreeFromDisk, fileTreeItemsFromDb)
	if syncError != nil {
		log.Fatal(syncError)
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

func syncFiles(databaseConnection *gorm.DB, filesFromDisk []models.FileTreeItem, filesFromDatabase []models.FileTreeItem) error {
	ctx := context.Background()
	filesToDelete := getDistinctFiles(filesFromDatabase, filesFromDisk)
	filesToCreate := getDistinctFiles(filesFromDisk, filesFromDatabase)

	deleteError := deleteFileTreeItemsFromDb(databaseConnection, ctx, filesToDelete)
	if deleteError != nil {
		return deleteError
	}

	insertError := insertFileTreeItemsIntoDb(databaseConnection, ctx, filesToCreate)
	if insertError != nil {
		return insertError
	}

	return nil
}

func getDistinctFiles(left []models.FileTreeItem, right []models.FileTreeItem) []models.FileTreeItem {
	distinctFiles := make([]models.FileTreeItem, 0)
	for _, leftItem := range left {
		isItemInBoothSets := false
		for _, rightItem := range right {
			if rightItem.Path == leftItem.Path {
				isItemInBoothSets = true
				continue
			}
		}

		if !isItemInBoothSets {
			distinctFiles = append(distinctFiles, leftItem)
		}
	}

	return distinctFiles
}

func deleteFileTreeItemsFromDb(databaseConnection *gorm.DB, ctx context.Context, filesToDelete []models.FileTreeItem) error {
	for _, fileToDelete := range filesToDelete {
		_, deleteError := gorm.G[models.FileTreeItem](databaseConnection).Where("id = ?", fileToDelete.ID).Delete(ctx)
		if deleteError != nil {
			log.Printf("Error while deleting file: %v\n", fileToDelete.Path)
			return deleteError
		}

		log.Printf("Successfully deleted file: %v\n", fileToDelete.Path)

		if fileToDelete.AssociatedAudioFileId != nil {
			doesAudioFileExits := doesFileWithFileIdExitsInDb(databaseConnection, *fileToDelete.AssociatedAudioFileId)
			if doesAudioFileExits {
				_, deleteError = gorm.G[models.FileTreeItem](databaseConnection).Where("file_id = ?", fileToDelete.AssociatedAudioFileId).Delete(ctx)
				if deleteError != nil {
					log.Printf("Error while deleting audio file with fileId='%v' of subtitle file with path='%v'\n", fileToDelete.AssociatedAudioFileId, fileToDelete.Path)
					return deleteError
				}

				log.Printf("Successfully deleted audio file with fileId='%v' of subtitle file with path='%v'\n", fileToDelete.AssociatedAudioFileId, fileToDelete.Path)
			}
		}

		if fileToDelete.LowQualityImageId != nil {
			doesLowQualityImageExist := doesFileWithFileIdExitsInDb(databaseConnection, *fileToDelete.LowQualityImageId)
			if doesLowQualityImageExist {
				_, deleteError = gorm.G[models.FileTreeItem](databaseConnection).Where("file_id = ?", fileToDelete.LowQualityImageId).Delete(ctx)
				if deleteError != nil {
					log.Printf("Error while deleting low quality image with fileId='%v' of item with path='%v'\n", fileToDelete.LowQualityImageId, fileToDelete.Path)
					return deleteError
				}

				log.Printf("Successfully deleted low quality image with fileId='%v' of item with path='%v'\n", fileToDelete.LowQualityImageId, fileToDelete.Path)
			}
		}
	}

	log.Printf("Deleted '%v' files\n", len(filesToDelete))
	return nil
}

func insertFileTreeItemsIntoDb(databaseConnection *gorm.DB, ctx context.Context, filesToAdd []models.FileTreeItem) error {
	// TODO this logic current can only create initial associations between items.
	// TODO adding subtitle files (associatedAudioFiles) or images (lowQualityImage) later on (e. G. first creating a mp3 file and later on creating a vtt file)
	// TODO will not be able to set corresponding associations (ids)
	// TODO there has to be added
	result := gorm.WithResult()
	createInBatchError := gorm.G[models.FileTreeItem](databaseConnection, result).CreateInBatches(ctx, &filesToAdd, len(filesToAdd))
	if createInBatchError != nil {
		log.Println("Error creating files")
		return createInBatchError
	}

	log.Printf("Created '%v' files\n", len(filesToAdd))
	return nil
}

func doesFileWithFileIdExitsInDb(databaseConnection *gorm.DB, fileId string) bool {
	count := int64(0)
	databaseConnection.Model(&models.FileTreeItem{}).Where("file_id = ?", fileId).Count(&count)
	return count > 0
}
