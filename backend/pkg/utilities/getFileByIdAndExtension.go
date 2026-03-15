package utilities

import (
	"backend/pkg/models"
	"path/filepath"
	"slices"
)

func GetFileByIdAndExtension(fileNodes []models.FileNode, id string, allowedExtension ...string) models.FileNode {
	var file models.FileNode
	for _, fileNode := range fileNodes {
		isMatch := fileNode.FileId == id
		if !isMatch {
			continue
		}

		if IsFileExtensionAllowed(fileNode, allowedExtension...) {
			file = fileNode
			break
		}
	}

	return file
}

func IsFileExtensionAllowed(fileNode models.FileNode, allowedExtension ...string) bool {
	ext := filepath.Ext(fileNode.Path)
	doesExtensionMatch := slices.Contains(allowedExtension, ext)
	return doesExtensionMatch
}
