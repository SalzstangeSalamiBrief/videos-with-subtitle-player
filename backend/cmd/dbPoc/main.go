package main

import (
	"backend/pkg/constants"
	"backend/pkg/enums/fileType"
	"backend/pkg/models"
	"backend/pkg/services/fileTreeManager"
	"backend/pkg/services/imageConverter/utilities"
	"context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"path"
	"strings"
)

// TODO CONFIGURATION
func main() {

	connectionString := "host=localhost port=5432 user=streamingServiceAccount password=tempPassword dbname=streaming-db sslmode=disable TimeZone=Europe/Berlin"
	databaseConnection, openDatabaseConnectionError := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if openDatabaseConnectionError != nil {
		log.Fatal(openDatabaseConnectionError)
	}

	ctx := context.Background()
	addFileTypeEnumError := getAndExecuteSqlFile(databaseConnection, ctx, "1_AddFileTypeEnum.sql")
	if addFileTypeEnumError != nil {
		log.Fatal(addFileTypeEnumError)
	}

	migrationError := databaseConnection.AutoMigrate(&models.FileTreeItem{}, &models.Tag{})
	if migrationError != nil {
		log.Fatal(migrationError)
	}

	initializedFileTreeManager := fileTreeManager.NewFileTreeManager("E:\\projects\\videos-with-subtitle-player\\test-content").InitializeTree()
	fileTreeFromDisk := initializedFileTreeManager.GetTree()

	seedTagsError := getAndExecuteSqlFile(databaseConnection, ctx, "2_tags_seed.sql")
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

func getAndExecuteSqlFile(db *gorm.DB, ctx context.Context, filename string) error {
	addFileTypeMigration, addFileTypeMigrationError := getMigrationFileContent(filename)
	if addFileTypeMigrationError != nil {
		return addFileTypeMigrationError
	}

	// TODO CREATE MIGRATION TABLE/EXECUTED SQL SCRIPTS
	// TODO CHECKSUM FOR ALL EXECUTED SCRIPTS

	executeAddFileTypeMigrationError := gorm.G[any](db).Exec(ctx, addFileTypeMigration)
	if executeAddFileTypeMigrationError != nil {
		return executeAddFileTypeMigrationError
	}

	return nil
}

// TODO CHECK PATHS ON EXECUTION
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
	if len(filesToDelete) == 0 {
		return nil
	}

	for _, fileToDelete := range filesToDelete {
		_, deleteError := gorm.G[models.FileTreeItem](databaseConnection).Where("id = ?", fileToDelete.ID).Delete(ctx)
		if deleteError != nil {
			log.Printf("Error while deleting file: %v\n", fileToDelete.Path)
			return deleteError
		}

		log.Printf("Successfully deleted file: %v\n", fileToDelete.Path)

		shouldTryCascadingDeletionOfAssociatedAudioFile := fileToDelete.AssociatedAudioFileId != nil
		if shouldTryCascadingDeletionOfAssociatedAudioFile {
			doesAudioFileExits := doesFileTreeItemWithFileIdExitsInDb(databaseConnection, *fileToDelete.AssociatedAudioFileId)
			if doesAudioFileExits {
				_, deleteError = gorm.G[models.FileTreeItem](databaseConnection).Where("file_id = ?", fileToDelete.AssociatedAudioFileId).Delete(ctx)
				if deleteError != nil {
					log.Printf("Error while deleting audio file with fileId='%v' of subtitle file with path='%v'\n", fileToDelete.AssociatedAudioFileId, fileToDelete.Path)
					return deleteError
				}

				log.Printf("Successfully deleted audio file with fileId='%v' of subtitle file with path='%v'\n", fileToDelete.AssociatedAudioFileId, fileToDelete.Path)
			}
		}

		shouldTryCascadingDeletionOfAssociatedLowQualityImage := fileToDelete.LowQualityImageId != nil
		if shouldTryCascadingDeletionOfAssociatedLowQualityImage {
			doesLowQualityImageExist := doesFileTreeItemWithFileIdExitsInDb(databaseConnection, *fileToDelete.LowQualityImageId)
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

func insertFileTreeItemsIntoDb(databaseConnection *gorm.DB, ctx context.Context, initialFilesToAdd []models.FileTreeItem) error {
	if len(initialFilesToAdd) == 0 {
		return nil
	}

	filesToAdd := make([]models.FileTreeItem, len(initialFilesToAdd))

	/**
	There are two cases files and their associated files can be added:
		1. On the same batch: The associated file is added in the same batch from the file system as the file, so it is part of the input => do nothing
		2. On a postposed batch: The associated file is not part of the same batch from the file system but could be already added to the database => try to get ids from the database

	Only these types of files have this problem:
		- subtitle
		- image
	*/
	for i, initialFileToAdd := range initialFilesToAdd {
		if initialFileToAdd.Type == fileType.IMAGE {
			isLowQuality := utilities.IsLowQualityImage(initialFileToAdd.Path)
			if isLowQuality {
				pathWithoutLowQualitySuffix := utilities.RemoveLowQualitySuffixFromImageName(initialFileToAdd.Path)

				isSameBatchInsert := false
				for _, referenceFile := range initialFilesToAdd {
					if referenceFile.Path == pathWithoutLowQualitySuffix {
						isSameBatchInsert = true
						break
					}
				}

				if !isSameBatchInsert {
					matchingImage, tryGetMatchingImageByPathError := tryGetFileTreeItemByPath(databaseConnection, ctx, pathWithoutLowQualitySuffix)
					if tryGetMatchingImageByPathError != nil {
						log.Printf("Warning: No matching audio file with path='%v' found. error='%v'\n", pathWithoutLowQualitySuffix, tryGetMatchingImageByPathError.Error())
					}

					if matchingImage != nil && matchingImage.ID > 0 {
						initialFileToAdd.LowQualityImageId = &matchingImage.FileId
					}
				}

			}
		}

		if initialFileToAdd.Type == fileType.SUBTITLE {
			pathOfPossibleMatchingAudioFile := strings.TrimSuffix(initialFileToAdd.Path, constants.SubtitleExtension)

			isSameBatchInsert := false
			for _, referenceFile := range initialFilesToAdd {
				if referenceFile.Path == pathOfPossibleMatchingAudioFile {
					isSameBatchInsert = true
					break
				}
			}

			if !isSameBatchInsert {
				matchingAudioFile, tryGetMatchingAudioFileByPathError := tryGetFileTreeItemByPath(databaseConnection, ctx, pathOfPossibleMatchingAudioFile)
				if tryGetMatchingAudioFileByPathError != nil {
					log.Printf("Warning: No matching audio file with path='%v' found. error='%v'\n", pathOfPossibleMatchingAudioFile, tryGetMatchingAudioFileByPathError.Error())
				}

				if matchingAudioFile != nil && matchingAudioFile.ID > 0 {
					initialFileToAdd.AssociatedAudioFileId = &matchingAudioFile.FileId
				}
			}
		}

		filesToAdd[i] = initialFileToAdd
	}

	result := gorm.WithResult()
	createInBatchError := gorm.G[models.FileTreeItem](databaseConnection, result).CreateInBatches(ctx, &filesToAdd, len(filesToAdd))
	if createInBatchError != nil {
		log.Println("Error creating files")
		return createInBatchError
	}

	return nil
}

func doesFileTreeItemWithFileIdExitsInDb(databaseConnection *gorm.DB, fileId string) bool {
	count := int64(0)
	databaseConnection.Model(&models.FileTreeItem{}).Where("file_id = ?", fileId).Count(&count)
	return count > 0
}

func tryGetFileTreeItemByPath(databaseConnection *gorm.DB, ctx context.Context, path string) (*models.FileTreeItem, error) {
	matchingFile, err := gorm.G[*models.FileTreeItem](databaseConnection).Where("path = ?", path).First(ctx)
	return matchingFile, err
}
