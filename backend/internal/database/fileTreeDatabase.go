package database

import (
	"backend/internal/configuration"
	"backend/pkg/models"
	"backend/pkg/services/fileTreeManager"
	"context"
	"errors"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type FileTreeDatabase struct {
	configuration      *configuration.DbConfiguration
	DatabaseConnection *gorm.DB
}

func NewFileTreeDatabase() *FileTreeDatabase {
	return &FileTreeDatabase{}
}

func (fileTree *FileTreeDatabase) AddConfiguration(configuration *configuration.DbConfiguration) *FileTreeDatabase {
	fileTree.configuration = configuration
	return fileTree
}

func (fileTreeDatabase *FileTreeDatabase) Open() (*FileTreeDatabase, error) {
	_, validationError := fileTreeDatabase.validateConfiguration()
	if validationError != nil {
		return nil, validationError
	}

	connectionString := fileTreeDatabase.configuration.GetConnectionString()
	databaseConnection, openDatabaseConnectionError := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if openDatabaseConnectionError != nil {
		log.Fatal(openDatabaseConnectionError)
	}

	fileTreeDatabase.DatabaseConnection = databaseConnection
	return fileTreeDatabase, nil
}

func (fileTreeDatabase *FileTreeDatabase) validateConfiguration() (*FileTreeDatabase, error) {
	if fileTreeDatabase.configuration == nil {
		return nil, errors.New("configuration is missing")
	}

	if fileTreeDatabase.configuration.Port <= 0 {
		return fileTreeDatabase, errors.New("configuration not complete: No 'Port' provided")
	}

	if fileTreeDatabase.configuration.Host == "" {
		return fileTreeDatabase, errors.New("configuration not complete: No 'Host' provided")
	}

	if fileTreeDatabase.configuration.Username == "" {
		return fileTreeDatabase, errors.New("configuration not complete: No 'Username' provided")
	}

	if fileTreeDatabase.configuration.Password == "" {
		return fileTreeDatabase, errors.New("configuration not complete: No 'Password' provided")
	}

	if fileTreeDatabase.configuration.DbName == "" {
		return fileTreeDatabase, errors.New("configuration not complete: No 'DbName' provided")
	}

	return fileTreeDatabase, nil
}

func (fileTreeDatabase *FileTreeDatabase) Close() error {
	if fileTreeDatabase.DatabaseConnection == nil {
		return errors.New("database connection not initialized")
	}

	db, err := fileTreeDatabase.DatabaseConnection.DB()
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}

	return nil
}

func (fileTreeDatabase *FileTreeDatabase) MigrateDatabase() (*FileTreeDatabase, error) {
	if fileTreeDatabase.DatabaseConnection == nil {
		return nil, errors.New("database connection not initialized")
	}

	ctx := context.Background()
	createMigrationsTableError := ExecuteMigration(fileTreeDatabase.DatabaseConnection, ctx, "0_CreateMigrationTable.sql", true)
	if createMigrationsTableError != nil {
		return fileTreeDatabase, createMigrationsTableError
	}

	addFileTypeEnumError := ExecuteMigration(fileTreeDatabase.DatabaseConnection, ctx, "1_AddFileTypeEnum.sql", false)
	if addFileTypeEnumError != nil {
		return fileTreeDatabase, addFileTypeEnumError
	}

	migrationError := fileTreeDatabase.DatabaseConnection.AutoMigrate(&models.FileTreeItem{}, &models.Tag{})
	if migrationError != nil {
		return fileTreeDatabase, migrationError
	}

	seedTagsError := ExecuteMigration(fileTreeDatabase.DatabaseConnection, ctx, "2_TagsSeed.sql", false)
	if seedTagsError != nil {
		log.Fatal(seedTagsError)
	}

	return fileTreeDatabase, nil
}

func (fileTreeDatabase *FileTreeDatabase) SyncFileTreeItems(manager *fileTreeManager.FileTreeManager) error {
	ctx := context.Background()

	fileTreeItemsFromDb, getFileTreeItemsFromDbError := gorm.G[models.FileTreeItem](fileTreeDatabase.DatabaseConnection).Find(ctx)
	if getFileTreeItemsFromDbError != nil {
		return getFileTreeItemsFromDbError
	}

	syncError := syncFileTree(fileTreeDatabase.DatabaseConnection, manager.GetTree(), fileTreeItemsFromDb, ctx)
	return syncError
}
