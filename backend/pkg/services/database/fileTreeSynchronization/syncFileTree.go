package fileTreeSynchronization

import (
	"backend/pkg/models"
	"backend/pkg/services/database/utilities"
	"context"
	"gorm.io/gorm"
)

func syncFileTree(databaseConnection *gorm.DB, filesFromDisk []models.FileTreeItem, filesFromDatabase []models.FileTreeItem) error {
	ctx := context.Background()
	filesToDelete := utilities.GetDistinctFiles(filesFromDatabase, filesFromDisk)
	filesToCreate := utilities.GetDistinctFiles(filesFromDisk, filesFromDatabase)

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
