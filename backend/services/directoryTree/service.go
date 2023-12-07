package directorytree

import (
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"videos-with-subtitle-player/models"

	"github.com/google/uuid"
)

var rootPath string

// TODO WHY SHOULD I SAVE FILE TREE; IF I WORK ONLY WITH FLAT TREE
var FileTree []models.DirectoryTreeItem
var FlatTree []models.FileTreeItem

func InitializeFileTree() {
	rootPath = os.Getenv("ROOT_PATH")
	fullTree := getFullTree(rootPath)
	cleanedTree := cleanUpTree(fullTree)
	FileTree = cleanedTree
	FlatTree = flattenTree(cleanedTree)
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
			Id:   uuid.New(),
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
				Id:   uuid.New(),
				Path: getFolderPath(currentItemPath),
				Name: getFileNameWithoutExtension(itemName),
			}

			newAssociatedSourceFile := models.FileTreeItem{
				Id:   uuid.New(),
				Path: getFolderPath(correspondingSourceFilePath),
				Name: getFileNameWithoutExtension(itemName),
			}

			newDirectoryItem.SubtitleFile = newSubtitleFile
			newDirectoryItem.AudioFile = newAssociatedSourceFile
		}

		currentFolderTree = append(currentFolderTree, newDirectoryItem)
	}

	return currentFolderTree
}
