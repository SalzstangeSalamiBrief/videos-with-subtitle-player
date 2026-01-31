package utilities

import "backend/pkg/models"

func GetDistinctFiles(left []models.FileTreeItem, right []models.FileTreeItem) []models.FileTreeItem {
	distinctFiles := make([]models.FileTreeItem, 0)
	for _, leftItem := range left {
		isItemInBoothSets := false
		for _, rightItem := range right {
			if rightItem.Path == leftItem.Path {
				isItemInBoothSets = true
				continue
			}
		}

		if !isItemInBoothSets {
			distinctFiles = append(distinctFiles, leftItem)
		}
	}

	return distinctFiles
}
