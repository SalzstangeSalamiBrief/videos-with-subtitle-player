package directorytree

import (
	"regexp"
	"strings"
	"videos-with-subtitle-player/models"
)

func cleanUpTree(originTree []models.DirectoryTreeItem) []models.DirectoryTreeItem {
	cleanedTree := make([]models.DirectoryTreeItem, 0)
	for _, item := range originTree {
		canBeAdded := checkIfDirectoryItemCanBeAdded(item)
		if canBeAdded == false {
			continue
		}
		item.Children = cleanUpTree(item.Children)
		item.Path = getFolderPath(item.Path)
		cleanedTree = append(cleanedTree, item)
	}

	return cleanedTree
}

func getFolderPath(path string) string {

	pathWithoutRoot := strings.Replace(path, rootPath, "", 1)
	regexpToAddmatchingSeparators := regexp.MustCompile(`\\+`)
	pathWithSeparators := regexpToAddmatchingSeparators.ReplaceAllString(pathWithoutRoot, "/")
	return pathWithSeparators
}

func getFileName(path string) string {
	regexpToAddmatchingSeparators := regexp.MustCompile(`\\+`)
	pathParts := regexpToAddmatchingSeparators.Split(path, -1)
	fileName := pathParts[len(pathParts)-1]
	return fileName
}

func checkIfDirectoryItemCanBeAdded(newDirectoryItem models.DirectoryTreeItem) bool {
	hasChildren := len(newDirectoryItem.Children) > 0
	hasAudioFile := newDirectoryItem.AudioFile.Path != ""
	hasSubtitleFile := newDirectoryItem.SubtitleFile.Path != ""
	canAdd := hasChildren || (hasAudioFile && hasSubtitleFile)
	return canAdd
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

func getFileNameWithoutExtension(filename string) string {
	// assume that there is no dot in the filename
	fileNameWithoutExtension := strings.Split(filename, ".")[0]
	return fileNameWithoutExtension
}
