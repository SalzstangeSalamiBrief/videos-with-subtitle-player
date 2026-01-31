package fileTreeSynchronization

import (
	"backend/pkg/models"
	"context"
	"gorm.io/gorm"
)

func tryGetFileTreeItemByPath(databaseConnection *gorm.DB, ctx context.Context, path string) (*models.FileTreeItem, error) {
	matchingFile, err := gorm.G[*models.FileTreeItem](databaseConnection).Where("path = ?", path).First(ctx)
	return matchingFile, err
}
