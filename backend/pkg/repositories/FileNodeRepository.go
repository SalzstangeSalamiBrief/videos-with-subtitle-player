package repositories

import (
	"backend/internal/database"
	"backend/pkg/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

type FileNodeRepository struct {
	database *database.Database
}

func NewFileTreeRepository(db *database.Database) (*FileNodeRepository, error) {
	if db == nil {
		return nil, errors.New("database is required")
	}

	return &FileNodeRepository{
		database: db,
	}, nil
}

func (repository *FileNodeRepository) GetFileNodes() ([]models.FileNode, error) {
	ctx := context.Background()
	fileNodesFromDb, getFileNodesFromDbError := gorm.G[models.FileNode](repository.database.DatabaseConnection).Find(ctx)
	return fileNodesFromDb, getFileNodesFromDbError
}

func (repository *FileNodeRepository) GetFileNodeById(fileId string) (models.FileNode, error) {
	ctx := context.Background()
	fileNodeFromDb, getFileNodeFromDbError := gorm.G[models.FileNode](repository.database.DatabaseConnection).Where("file_id = ?", fileId).First(ctx)
	return fileNodeFromDb, getFileNodeFromDbError
}
