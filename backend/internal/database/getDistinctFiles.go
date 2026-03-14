package database

import "backend/pkg/models"

func getDistinctFiles(left []models.FileNode, right []models.FileNode) []models.FileNode {
	distinctFiles := make([]models.FileNode, 0)
	for _, leftItem := range left {
		isItemInBoothSets := false
		for _, rightItem := range right {
			// TODO DOES THIS WORK? CURRENTLY IDSS GETTING OVERWRITTEN
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
