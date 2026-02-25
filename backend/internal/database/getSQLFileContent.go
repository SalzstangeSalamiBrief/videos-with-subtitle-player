package utilities

import (
	"embed"
)

//go:embed ../migrations/*.sql
var fileTreeDatabaseMigrationFiles embed.FS

// TODO CHECK PATHS ON EXECUTION
func GetSQLFileContent(filename string) (string, error) {
	contentAsByteArray, readFileError := fileTreeDatabaseMigrationFiles.ReadFile("migrations" + filename)
	//var relativeFilePathToMigrationFolder = path.Join("./migrations/", filename)
	//contentAsByteArray, readFileError := os.ReadFile(relativeFilePathToMigrationFolder)
	if readFileError != nil {
		return "", readFileError
	}

	var stringifiedContent = string(contentAsByteArray)
	return stringifiedContent, nil
}
