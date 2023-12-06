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

var fileTree []models.DirectoryTreeItem
var rootPath string

func GetFileTreeUseCase(w http.ResponseWriter, r *http.Request, quit chan<- bool) {
	if rootPath == "" {
		rootPath = os.Getenv("ROOT_PATH")
	}

	if len(fileTree) == 0 {
		fileTree = getFileTree(rootPath)
	} else {
		// update file tree in background
		go func() {
			fileTree = getFileTree(rootPath)
		}()
	}

	encodedBytes, err := json.Marshal(fileTree)
	if err != nil {
		router.ErrorHandler(w, fmt.Sprintf("Could not marshal file tree: %v", err.Error()), http.StatusInternalServerError)
		quit <- true
		return
	}

	w.Write(encodedBytes)
	quit <- true
}

func getFileTree(parentPath string) []models.DirectoryTreeItem {
	fullTree := getFullTree(parentPath)
	cleanedTree := removeEmptyDirectoryItems(fullTree)

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

			fmt.Printf("fileExtension %v\n", fileExtension)
			newSubtitleFile := models.FileTreeItem{
				Path: currentItemPath,
				Name: itemName,
			}

			newAssociatedSourceFile := models.FileTreeItem{
				Path: parentPath + "\\" + correspondingSourceFilePath,
				Name: itemName,
			}

			newDirectoryItem.SubtitleFile = newSubtitleFile
			newDirectoryItem.AudioFile = newAssociatedSourceFile
		}

		currentFolderTree = append(currentFolderTree, newDirectoryItem)
	}

	return currentFolderTree
}

func removeEmptyDirectoryItems(originTree []models.DirectoryTreeItem) []models.DirectoryTreeItem {
	cleanedTree := make([]models.DirectoryTreeItem, 0)
	for _, item := range originTree {
		canBeAdded := checkIfDirectoryItemCanBeAdded(item)
		if canBeAdded == true {
			item.Children = removeEmptyDirectoryItems(item.Children)
			cleanedTree = append(cleanedTree, item)
		}
	}

	return cleanedTree
}

func checkIfDirectoryItemCanBeAdded(newDirectoryItem models.DirectoryTreeItem) bool {
	hasChildren := len(newDirectoryItem.Children) > 0
	hasAudioFile := newDirectoryItem.AudioFile.Path != ""
	hasSubtitleFile := newDirectoryItem.SubtitleFile.Path != ""
	canAdd := hasChildren || (hasAudioFile && hasSubtitleFile)
	return canAdd
}
