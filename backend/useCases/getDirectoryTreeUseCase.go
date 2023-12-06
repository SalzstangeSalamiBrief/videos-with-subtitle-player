package usecases

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"videos-with-subtitle-player/models"
	"videos-with-subtitle-player/router"
)

var GetFileTreeUseCaseRoute = router.Route{
	Path:    "/file-tree",
	Handler: GetFileTreeUseCase,
	Method:  http.MethodGet,
}

var rootPath string

func GetFileTreeUseCase(w http.ResponseWriter, r *http.Request, quit chan<- bool) {
	if rootPath == "" {
		rootPath = os.Getenv("ROOT_PATH")
	}

	fileTree := GetFileTree(rootPath)
	// TODO CACHING
	// TODO ACCESS AND LOAD FILE TREE FROM OTHJER FILES

	encodedBytes, err := json.Marshal(fileTree)
	if err != nil {
		router.ErrorHandler(w, fmt.Sprintf("Could not marshal file tree: %v", err.Error()), http.StatusInternalServerError)
		quit <- true
		return
	}

	w.Write(encodedBytes)
	quit <- true
}

func GetFileTree(parentPath string) []models.DirectoryTreeItem {
	fullTree := getFullTree(parentPath)
	cleanedTree := cleanUpTree(fullTree)

	return cleanedTree
}

func getFullTree(parentPath string) []models.DirectoryTreeItem {
	content, err := os.ReadDir(parentPath)
	if err != nil {
		log.Fatal(err)
	}

	currentFolderTree := make([]models.DirectoryTreeItem, 0)

	for _, item := range content {
		itemName := item.Name()

		currentItemPath := parentPath + "\\" + itemName
		isDirectory := item.IsDir()

		newDirectoryItem := models.DirectoryTreeItem{
			Path: currentItemPath,
			Name: itemName,
		}

		if isDirectory == true {
			children := getFullTree(currentItemPath)
			newDirectoryItem.Children = children
			currentFolderTree = append(currentFolderTree, newDirectoryItem)
		}

		if isDirectory == false {
			fileExtension := path.Ext(itemName)
			isSubtitleFile := fileExtension == ".vtt"

			if isSubtitleFile == false {
				continue
			}

			mp3RegExp := regexp.MustCompile(".mp3.vtt$")
			isAssociatedWithMp3File := mp3RegExp.MatchString(itemName)

			if isAssociatedWithMp3File == false {
				log.Default().Printf("'%v' is not associated with an mp3 file\n", itemName)
				continue
			}

			correspondingSourceFilePath := strings.Replace(currentItemPath, fileExtension, "", 1)
			_, err := os.Stat(correspondingSourceFilePath)
			if os.IsNotExist(err) {
				log.Default().Printf("corresponding source file does not exist for '%v'\n", itemName)
				continue
			}

			if err != nil {
				log.Default().Printf("error while checking if corresponding source file exists for '%v': '%v'\n", itemName, err.Error())
				continue
			}

			newSubtitleFile := models.FileTreeItem{
				Path: getFolderPath(currentItemPath),
				Name: itemName,
			}

			newAssociatedSourceFile := models.FileTreeItem{
				Path: getFolderPath(correspondingSourceFilePath),
				Name: itemName,
			}

			newDirectoryItem.SubtitleFile = newSubtitleFile
			newDirectoryItem.AudioFile = newAssociatedSourceFile
		}

		currentFolderTree = append(currentFolderTree, newDirectoryItem)
	}

	return currentFolderTree
}

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
	if rootPath == "" {
		rootPath = os.Getenv("ROOT_PATH")
	}

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
