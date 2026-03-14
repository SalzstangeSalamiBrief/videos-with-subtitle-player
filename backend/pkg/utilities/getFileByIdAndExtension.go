package utilities

import (
	"backend/pkg/models"
	"path/filepath"
	"slices"
)

func GetFileByIdAndExtension(fileTreeItems []models.FileNode, id string, allowedExtension ...string) models.FileNode {
	var file models.FileNode
	for _, fileTreeItem := range fileTreeItems {
		isMatch := fileTreeItem.FileId == id
		if !isMatch {
			continue
		}

		if IsFileExtensionAllowed(fileTreeItem, allowedExtension...) {
			file = fileTreeItem
			break
		}
	}

	return file
}

func IsFileExtensionAllowed(fileTreeItem models.FileNode, allowedExtension ...string) bool {
	ext := filepath.Ext(fileTreeItem.Path)
	doesExtensionMatch := slices.Contains(allowedExtension, ext)
	return doesExtensionMatch
}
