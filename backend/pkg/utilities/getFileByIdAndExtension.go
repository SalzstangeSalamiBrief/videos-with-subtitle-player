package utilities

import (
	"backend/pkg/models"
	"path/filepath"
	"slices"
)

func GetFileByIdAndExtension(fileTreeItems []models.FileTreeItem, id string, allowedExtension ...string) models.FileTreeItem {
	var file models.FileTreeItem
	for _, fileTreeItem := range fileTreeItems {
		isMatch := fileTreeItem.FileId == id
		if !isMatch {
			continue
		}

		if isFileExtensionAllowed(fileTreeItem, allowedExtension...) {
			file = fileTreeItem
			break
		}
	}

	return file
}

func isFileExtensionAllowed(fileTreeItem models.FileTreeItem, allowedExtension ...string) bool {
	ext := filepath.Ext(fileTreeItem.Path)
	doesExtensionMatch := slices.Contains(allowedExtension, ext)
	return doesExtensionMatch
}
