package database

import (
	"backend/internal/configuration"
	"backend/pkg/models"
	"backend/pkg/services/fileTreeManager"
	"context"
	"errors"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	configuration      *configuration.DbConfiguration
	DatabaseConnection *gorm.DB
}

func NewDatabase() *Database {
	return &Database{}
}

func (db *Database) AddConfiguration(configuration *configuration.DbConfiguration) *Database {
	db.configuration = configuration
	return db
}

func (db *Database) Open() (*Database, error) {
	_, validationError := db.validateConfiguration()
	if validationError != nil {
		return nil, validationError
	}

	connectionString := db.configuration.GetConnectionString()
	databaseConnection, openDatabaseConnectionError := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if openDatabaseConnectionError != nil {
		log.Fatal(openDatabaseConnectionError)
	}

	sqlDB, err := databaseConnection.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	db.DatabaseConnection = databaseConnection
	return db, nil
}

func (db *Database) validateConfiguration() (*Database, error) {
	if db.configuration == nil {
		return nil, errors.New("configuration is missing")
	}

	if db.configuration.Port <= 0 {
		return db, errors.New("configuration not complete: No 'Port' provided")
	}

	if db.configuration.Host == "" {
		return db, errors.New("configuration not complete: No 'Host' provided")
	}

	if db.configuration.Username == "" {
		return db, errors.New("configuration not complete: No 'Username' provided")
	}

	if db.configuration.Password == "" {
		return db, errors.New("configuration not complete: No 'Password' provided")
	}

	if db.configuration.DbName == "" {
		return db, errors.New("configuration not complete: No 'DbName' provided")
	}

	return db, nil
}

func (db *Database) Close() error {
	if db.DatabaseConnection == nil {
		return errors.New("database connection not initialized")
	}

	openedDb, err := db.DatabaseConnection.DB()
	if err != nil {
		return err
	}

	err = openedDb.Close()
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) MigrateDatabase() (*Database, error) {
	if db.DatabaseConnection == nil {
		return nil, errors.New("database connection not initialized")
	}

	ctx := context.Background()
	createMigrationsTableError := executeMigration(db.DatabaseConnection, ctx, "0_CreateMigrationTable.sql", true)
	if createMigrationsTableError != nil {
		return db, createMigrationsTableError
	}

	addFileTypeEnumError := executeMigration(db.DatabaseConnection, ctx, "1_AddFileTypeEnum.sql", false)
	if addFileTypeEnumError != nil {
		return db, addFileTypeEnumError
	}

	migrationError := db.DatabaseConnection.AutoMigrate(&models.FileTreeItem{}, &models.Tag{})
	if migrationError != nil {
		return db, migrationError
	}

	seedTagsError := executeMigration(db.DatabaseConnection, ctx, "2_TagsSeed.sql", false)
	if seedTagsError != nil {
		log.Fatal(seedTagsError)
	}

	return db, nil
}

func (db *Database) SyncFileTreeItems(manager *fileTreeManager.FileTreeManager) error {
	ctx := context.Background()

	fileTreeItemsFromDb, getFileTreeItemsFromDbError := gorm.G[models.FileTreeItem](db.DatabaseConnection).Find(ctx)
	if getFileTreeItemsFromDbError != nil {
		return getFileTreeItemsFromDbError
	}

	syncError := syncFileTree(db.DatabaseConnection, manager.GetTree(), fileTreeItemsFromDb, ctx)
	return syncError
}
