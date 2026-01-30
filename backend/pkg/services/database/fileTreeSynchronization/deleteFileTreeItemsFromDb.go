package fileTreeSynchronization

import (
	"backend/pkg/models"
	"context"
	"gorm.io/gorm"
	"log"
)

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
