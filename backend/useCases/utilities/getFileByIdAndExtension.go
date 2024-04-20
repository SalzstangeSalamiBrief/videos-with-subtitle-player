package usecases

import (
	"backend/lib"
	"backend/models"
	"path/filepath"
	"slices"
)

func GetFileByIdAndExtension(id string, allowedExtension ...string) models.FileTreeItem {
	var file models.FileTreeItem
	for _, fileTreeItem := range lib.FileTreeItems {
		isMatch := fileTreeItem.Id == id
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
