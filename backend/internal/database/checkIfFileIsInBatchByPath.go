package database

import "backend/pkg/models"

func checkIfFileIsInBatchByPath(path string, batch []models.FileNode) bool {
	doesFileWithPathExist := false
	for _, referenceFile := range batch {
		if referenceFile.Path == path {
			doesFileWithPathExist = true
			break
		}
	}

	return doesFileWithPathExist
}
