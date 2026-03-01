package database

import "backend/pkg/models"

func checkIfFileIsInBatchByPath(path string, batch []models.FileTreeItem) bool {
	doesFileWithPathExist := false
	for _, referenceFile := range batch {
		if referenceFile.Path == path {
			doesFileWithPathExist = true
			break
		}
	}

	return doesFileWithPathExist
}
