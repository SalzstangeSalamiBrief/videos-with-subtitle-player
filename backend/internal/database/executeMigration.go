package database

import (
	"backend/internal/database/models"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func executeMigration(db *gorm.DB, ctx context.Context, filename string, skipValidation bool) error {
	fileContent, getFileContentError := getSQLFileContent(filename)
	if getFileContentError != nil {
		return getFileContentError
	}

	if !skipValidation {
		canExecute, canExecuteError := canExecuteMigration(db, ctx, filename, fileContent)
		if canExecuteError != nil {
			return canExecuteError
		}

		if !canExecute {
			fmt.Printf("Cannot execute migration for file '%s'\n", filename)
			return nil
		}
	}

	executeAddFileTypeMigrationError := gorm.G[any](db).Exec(ctx, fileContent)
	if executeAddFileTypeMigrationError != nil {
		return executeAddFileTypeMigrationError
	}

	createMigrationEntryError := createMigrationEntry(db, ctx, filename, fileContent)
	if createMigrationEntryError != nil {
		return createMigrationEntryError
	}

	return nil
}

func canExecuteMigration(db *gorm.DB, ctx context.Context, filename string, fileContent string) (bool, error) {
	receivedMigration, err := gorm.G[models.Migration](db).Where("name = ?", filename).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, nil
		}

		return false, err
	}

	expectedChecksum := createChecksumFromString(fileContent)
	if receivedMigration.Checksum != expectedChecksum {
		return false, errors.New(fmt.Sprintf("Mismatch of checksum for item with name='%s'. Expected checksum '%s' but received '%s'", filename, expectedChecksum, receivedMigration.Checksum))
	}

	return true, nil
}

func createMigrationEntry(db *gorm.DB, ctx context.Context, filename string, fileContent string) error {
	_, doesMigrationExistsError := gorm.G[models.Migration](db).Where("name = ?", filename).First(ctx)
	if doesMigrationExistsError == nil {
		return nil
	}

	migrationToCreate := models.Migration{
		Name:     filename,
		Checksum: createChecksumFromString(fileContent),
	}
	err := gorm.G[models.Migration](db).Create(ctx, &migrationToCreate)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil
		}

		return err
	}

	return nil
}

func createChecksumFromString(input string) string {
	bytes := sha256.Sum256([]byte(input))
	return fmt.Sprintf("%x", bytes)
}
