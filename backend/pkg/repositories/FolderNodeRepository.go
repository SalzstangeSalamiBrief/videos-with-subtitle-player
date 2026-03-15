package repositories

import (
	"backend/internal/database"
	"backend/pkg/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

type FolderNodeRepository struct {
	database *database.Database
}

func NewFolderNodeRepository(db *database.Database) (*FolderNodeRepository, error) {
	if db == nil {
		return nil, errors.New("database is required")
	}

	return &FolderNodeRepository{
		database: db,
	}, nil
}

func (repository *FolderNodeRepository) GetFolders() ([]models.FolderNode, error) {
	ctx := context.Background()
	folderNodesFromDb, getFolderNodesFromDbError := gorm.G[models.FolderNode](repository.database.DatabaseConnection).Find(ctx)
	return folderNodesFromDb, getFolderNodesFromDbError
}
