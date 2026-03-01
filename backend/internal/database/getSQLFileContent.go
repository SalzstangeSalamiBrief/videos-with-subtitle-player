package database

import (
	"embed"
)

//go:embed migrations/*.sql
var fileTreeDatabaseMigrationFiles embed.FS

func getSQLFileContent(filename string) (string, error) {
	contentAsByteArray, readFileError := fileTreeDatabaseMigrationFiles.ReadFile("migrations/" + filename)
	if readFileError != nil {
		return "", readFileError
	}

	var stringifiedContent = string(contentAsByteArray)
	return stringifiedContent, nil
}
