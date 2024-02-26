package directorytree

import (
	"backend/models"
	"github.com/google/uuid"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

var rootPath string
var FileTreeItems []models.FileTreeItem

func InitializeFileTree() {
	rootPath = os.Getenv("ROOT_PATH")
	fullTree := getFullTree(rootPath)
	FileTreeItems = fullTree
}

func getFullTree(parentPath string) []models.FileTreeItem {
	content, err := os.ReadDir(parentPath)
	if err != nil {
		log.Fatal(err)
	}

	currentFileItems := make([]models.FileTreeItem, 0)

	for _, item := range content {
		itemName := item.Name()
		currentItemPath := parentPath + "\\" + itemName
		isDirectory := item.IsDir()

		if isDirectory {
			newDirectoryItems := getFullTree(currentItemPath)
			currentFileItems = append(currentFileItems, newDirectoryItems...)
			continue
		}

		fileExtension := path.Ext(itemName)
		isSubtitleFile := fileExtension == ".vtt"

		if !isSubtitleFile {
			continue
		}

		mp3RegExp := regexp.MustCompile(".mp3.vtt$")
		isAssociatedWithMp3File := mp3RegExp.MatchString(itemName)

		if !isAssociatedWithMp3File {
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
			Id:   uuid.New().String(),
			Path: getFolderPath(currentItemPath),
			Name: getFileNameWithoutExtension(itemName),
		}

		newAssociatedSourceFile := models.FileTreeItem{
			Id:   uuid.New().String(),
			Path: getFolderPath(correspondingSourceFilePath),
			Name: getFileNameWithoutExtension(itemName),
		}

		currentFileItems = append(currentFileItems, newSubtitleFile)
		currentFileItems = append(currentFileItems, newAssociatedSourceFile)
	}

	return currentFileItems
}
