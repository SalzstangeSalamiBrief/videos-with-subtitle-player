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
var FileTreeItems []models.FileTreeItem

func InitializeFileTree() {
	rootPath = os.Getenv("ROOT_PATH")
	fullTree := getFullTree(rootPath)
	flattenedTree := fullTree
	FileTreeItems = flattenedTree
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

		if isDirectory == true {
			newDirectoryItems := getFullTree(currentItemPath)
			currentFileItems = append(currentFileItems, newDirectoryItems...)
			continue
		}

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

		currentFileItems = append(currentFileItems, newSubtitleFile)
		currentFileItems = append(currentFileItems, newAssociatedSourceFile)
	}

	return currentFileItems
}
