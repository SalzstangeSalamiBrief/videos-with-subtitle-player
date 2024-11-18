package utilities

import (
	"backend/pkg/models"
	"backend/pkg/services/fileTreeManager"
)

func GetFileByIdAndExtension(id string, allowedExtension ...string) models.FileTreeNode {
	var file models.FileTreeNode
	for _, fileTreeItem := range fileTreeManager.FileTreeNodes {
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
