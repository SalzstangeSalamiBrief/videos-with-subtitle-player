package database

import (
	"backend/pkg/constants"
	"backend/pkg/enums/fileType"
	"backend/pkg/models"
	imageConverterUtilities "backend/pkg/services/imageConverter/utilities"
	"context"
	"log"
	"strings"

	"gorm.io/gorm"
)

func syncFiles(databaseConnection *gorm.DB, filesFromDisk []models.FileNode, filesFromDatabase []models.FileNode, ctx context.Context) error {
	filesToDelete := getDistinctFiles(filesFromDatabase, filesFromDisk)
	filesToCreate := getDistinctFiles(filesFromDisk, filesFromDatabase)
	deleteError := deleteFileNodesFromDb(databaseConnection, ctx, filesToDelete)
	if deleteError != nil {
		return deleteError
	}

	insertError := insertFileNodessIntoDb(databaseConnection, ctx, filesToCreate)
	if insertError != nil {
		return insertError
	}

	return nil
}

func doesFileNodeWithFileIdExitsInDb(databaseConnection *gorm.DB, fileId string) bool {
	count := int64(0)
	databaseConnection.Model(&models.FileNode{}).Where("file_id = ?", fileId).Count(&count)
	return count > 0
}

func deleteFileNodesFromDb(databaseConnection *gorm.DB, ctx context.Context, filesToDelete []models.FileNode) error {
	if len(filesToDelete) == 0 {
		return nil
	}

	for _, fileToDelete := range filesToDelete {
		_, deleteError := gorm.G[models.FileNode](databaseConnection).Where("id = ?", fileToDelete.ID).Delete(ctx)
		if deleteError != nil {
			log.Printf("Error while deleting file: %v\n", fileToDelete.Path)
			return deleteError
		}

		log.Printf("Successfully deleted file: %v\n", fileToDelete.Path)

		shouldTryCascadingDeletionOfAssociatedAudioFile := fileToDelete.AssociatedAudioFileId != nil
		if shouldTryCascadingDeletionOfAssociatedAudioFile {
			doesAudioFileExits := doesFileNodeWithFileIdExitsInDb(databaseConnection, *fileToDelete.AssociatedAudioFileId)
			if doesAudioFileExits {
				_, deleteError = gorm.G[models.FileNode](databaseConnection).Where("file_id = ?", fileToDelete.AssociatedAudioFileId).Delete(ctx)
				if deleteError != nil {
					log.Printf("Error while deleting audio file with fileId='%v' of subtitle file with path='%v'\n", fileToDelete.AssociatedAudioFileId, fileToDelete.Path)
					return deleteError
				}

				log.Printf("Successfully deleted audio file with fileId='%v' of subtitle file with path='%v'\n", fileToDelete.AssociatedAudioFileId, fileToDelete.Path)
			}
		}

		shouldTryCascadingDeletionOfAssociatedLowQualityImage := fileToDelete.LowQualityImageId != nil
		if shouldTryCascadingDeletionOfAssociatedLowQualityImage {
			doesLowQualityImageExist := doesFileNodeWithFileIdExitsInDb(databaseConnection, *fileToDelete.LowQualityImageId)
			if doesLowQualityImageExist {
				_, deleteError = gorm.G[models.FileNode](databaseConnection).Where("file_id = ?", fileToDelete.LowQualityImageId).Delete(ctx)
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

func insertFileNodessIntoDb(databaseConnection *gorm.DB, ctx context.Context, filesToAddInput []models.FileNode) error {
	if len(filesToAddInput) == 0 {
		return nil
	}

	filesToAdd := make([]models.FileNode, len(filesToAddInput))

	/**
	There are two cases files and their associated files can be added:
		1. On the same batch: The associated file is added in the same batch from the file system as the file, so it is part of the input => do nothing
		2. On a postposed batch: The associated file is not part of the same batch from the file system but could be already added to the database => try to get ids from the database

	Only these types of files have this problem:
		- subtitle
		- image
	*/
	for i, initialFileToAdd := range filesToAddInput {
		if initialFileToAdd.Type == fileType.IMAGE {
			isLowQuality := imageConverterUtilities.IsLowQualityImagePath(initialFileToAdd.Path)
			if isLowQuality {
				pathWithoutLowQualitySuffix := imageConverterUtilities.RemoveLowQualitySuffixFromImageName(initialFileToAdd.Path)
				isSameBatchInsert := checkIfFileIsInBatchByPath(pathWithoutLowQualitySuffix, filesToAddInput)

				if !isSameBatchInsert {
					matchingImage, tryGetMatchingImageByPathError := tryGetFileNodeByPath(databaseConnection, ctx, pathWithoutLowQualitySuffix)
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
			isSameBatchInsert := checkIfFileIsInBatchByPath(pathOfPossibleMatchingAudioFile, filesToAddInput)

			if !isSameBatchInsert {
				matchingAudioFile, tryGetMatchingAudioFileByPathError := tryGetFileNodeByPath(databaseConnection, ctx, pathOfPossibleMatchingAudioFile)
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
	createInBatchError := gorm.G[models.FileNode](databaseConnection, result).CreateInBatches(ctx, &filesToAdd, len(filesToAdd))
	if createInBatchError != nil {
		log.Println("Error creating files")
		return createInBatchError
	}

	return nil
}

func tryGetFileNodeByPath(databaseConnection *gorm.DB, ctx context.Context, path string) (*models.FileNode, error) {
	matchingFile, err := gorm.G[*models.FileNode](databaseConnection).Where("path = ?", path).First(ctx)
	return matchingFile, err
}
