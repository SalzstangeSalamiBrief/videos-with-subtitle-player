package fileTreeSynchronization

import (
	"backend/pkg/models"
	"gorm.io/gorm"
)

func doesFileTreeItemWithFileIdExitsInDb(databaseConnection *gorm.DB, fileId string) bool {
	count := int64(0)
	databaseConnection.Model(&models.FileTreeItem{}).Where("file_id = ?", fileId).Count(&count)
	return count > 0
}
