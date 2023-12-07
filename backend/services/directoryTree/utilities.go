package directorytree

import (
	"path"
	"regexp"
	"strings"
	"videos-with-subtitle-player/models"
)

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
		// Append audio file if it exists
		if item.AudioFile.Path != "" {
			flatTree = appendIfNotExists(flatTree, item.AudioFile)
		}

		// Append subtitle file if it exists
		if item.SubtitleFile.Path != "" {
			flatTree = appendIfNotExists(flatTree, item.SubtitleFile)
		}

		// Recursively flatten children
		flatChildren := flattenTree(item.Children)
		flatTree = append(flatTree, flatChildren...)
	}

	return flatTree
}

func appendIfNotExists(slice []models.FileTreeItem, item models.FileTreeItem) []models.FileTreeItem {
	for _, existingItem := range slice {
		if existingItem.Path == item.Path {
			return slice
		}
	}
	return append(slice, item)
}

func getFileNameWithoutExtension(filename string) string {
	fileExtension := path.Ext(filename)
	fileNameWithoutExtension := strings.Replace(filename, fileExtension, "", 1)
	// used if two file names are chained
	fileExtension = path.Ext(fileNameWithoutExtension)
	if fileExtension != "" {
		fileNameWithoutExtension = strings.Replace(fileNameWithoutExtension, fileExtension, "", 1)
	}

	return fileNameWithoutExtension
}
