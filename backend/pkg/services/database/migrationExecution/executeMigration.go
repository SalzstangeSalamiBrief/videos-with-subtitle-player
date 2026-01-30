package migrationExecution

import (
	"backend/pkg/services/database/utilities"
	"context"
	"gorm.io/gorm"
)

func ExecuteMigration(db *gorm.DB, ctx context.Context, filename string) error {
	addFileTypeMigration, addFileTypeMigrationError := utilities.GetSQLFileContent(filename)
	if addFileTypeMigrationError != nil {
		return addFileTypeMigrationError
	}

	// TODO CREATE MIGRATION TABLE/EXECUTED SQL SCRIPTS
	// TODO CHECKSUM FOR ALL EXECUTED SCRIPTS

	executeAddFileTypeMigrationError := gorm.G[any](db).Exec(ctx, addFileTypeMigration)
	if executeAddFileTypeMigrationError != nil {
		return executeAddFileTypeMigrationError
	}

	return nil
}
