package utilities

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

		ext := filepath.Ext(fileTreeItem.Path)
		doesExtensionMatch := slices.Contains(allowedExtension, ext)
		if doesExtensionMatch {
			file = fileTreeItem
			break
		}
	}

	return file
}
