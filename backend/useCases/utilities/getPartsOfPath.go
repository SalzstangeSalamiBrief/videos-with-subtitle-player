package usecases

import (
	"backend/models"
	"path/filepath"
	"strings"
)

func GetPartsOfPath(file models.FileTreeItem) []string {
	filePath, _ := filepath.Split(file.Path)
	allParts := strings.Split(filePath, string(filepath.Separator))
	var parts []string
	for _, part := range allParts {
		if part == "" {
			continue
		}

		parts = append(parts, part)
	}

	return parts
}
