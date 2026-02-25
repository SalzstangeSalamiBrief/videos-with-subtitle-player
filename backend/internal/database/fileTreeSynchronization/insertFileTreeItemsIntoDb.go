package fileTreeSynchronization

import (
	"backend/internal/database"
	"backend/pkg/constants"
	"backend/pkg/enums/fileType"
	"backend/pkg/models"
	imageConverterUtilities "backend/pkg/services/imageConverter/utilities"
	"context"
	"log"
	"strings"

	"gorm.io/gorm"
)

func insertFileTreeItemsIntoDb(databaseConnection *gorm.DB, ctx context.Context, filesToAddInput []models.FileTreeItem) error {
	if len(filesToAddInput) == 0 {
		return nil
	}

	filesToAdd := make([]models.FileTreeItem, len(filesToAddInput))

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
				isSameBatchInsert := database.CheckIfFileIsInBatchByPath(pathWithoutLowQualitySuffix, filesToAddInput)

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
			isSameBatchInsert := database.CheckIfFileIsInBatchByPath(pathOfPossibleMatchingAudioFile, filesToAddInput)

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
