package usecases

import "videos-with-subtitle-player/models"

func GetFlatFileTree(parentPath string) []models.FileTreeItem {
	fullTree := GetFileTree(parentPath)
	flatTree := flattenTree(fullTree)
	return flatTree
}

func flattenTree(originTree []models.DirectoryTreeItem) []models.FileTreeItem {
	flatTree := make([]models.FileTreeItem, 0)
	for _, item := range originTree {
		hasAudioFile := item.AudioFile.Path != ""
		hasSubtitleFile := item.SubtitleFile.Path != ""
		if hasAudioFile && hasSubtitleFile {
			flatTree = append(flatTree, item.AudioFile)
			flatTree = append(flatTree, item.SubtitleFile)
			continue
		}

		flatChildren := flattenTree(item.Children)
		flatTree = append(flatTree, flatChildren...)
	}

	return flatTree
}
