package utilities

import (
	"backend/pkg/models"
	"backend/pkg/services/fileTreeManager"
)

func GetFileByIdAndExtension(id string, allowedExtension ...string) models.FileTreeItem {
	var file models.FileTreeItem
	for _, fileTreeItem := range fileTreeManager.FileTreeItems {
		isMatch := fileTreeItem.Id == id
		if !isMatch {
			continue
		}

		if fileTreeItem.IsFileExtensionAllowed(allowedExtension...) {
			file = fileTreeItem
			break
		}
	}

	return file
}
