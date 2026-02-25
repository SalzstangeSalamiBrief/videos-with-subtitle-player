package fileTreeSynchronization

import (
	"backend/internal/database"
	"backend/pkg/models"
	"context"

	"gorm.io/gorm"
)

func syncFileTree(databaseConnection *gorm.DB, filesFromDisk []models.FileTreeItem, filesFromDatabase []models.FileTreeItem) error {
	ctx := context.Background()
	filesToDelete := database.GetDistinctFiles(filesFromDatabase, filesFromDisk)
	filesToCreate := database.GetDistinctFiles(filesFromDisk, filesFromDatabase)

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
