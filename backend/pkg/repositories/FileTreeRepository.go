package repositories

import (
	"backend/internal/database"
	"backend/pkg/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

type FileTreeRepository struct {
	database *database.Database
}

func NewFileTreeRepository(db *database.Database) (*FileTreeRepository, error) {
	if db == nil {
		return nil, errors.New("database is required")
	}

	return &FileTreeRepository{
		database: db,
	}, nil
}

func (repository *FileTreeRepository) GetFileTree() ([]models.FileTreeItem, error) {
	ctx := context.Background()
	fileTreeItemsFromDb, getFileTreeItemsFromDbError := gorm.G[models.FileTreeItem](repository.database.DatabaseConnection).Find(ctx)
	return fileTreeItemsFromDb, getFileTreeItemsFromDbError
}

func (repository *FileTreeRepository) GetFileByFileId(fileId string) (models.FileTreeItem, error) {
	ctx := context.Background()
	fileTreeItemsFromDb, getFileTreeItemsFromDbError := gorm.G[models.FileTreeItem](repository.database.DatabaseConnection).Where("file_id = ?", fileId).First(ctx)
	return fileTreeItemsFromDb, getFileTreeItemsFromDbError
}
