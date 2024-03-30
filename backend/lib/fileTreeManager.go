package lib

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

		log.Default().Printf("'%v' is a directory\n", itemName)
		if isDirectory {
			newDirectoryItems := getFullTree(currentItemPath)
			currentFileItems = append(currentFileItems, newDirectoryItems...)
			continue
		}

		fileExtension := path.Ext(itemName)

		isVideoFile := fileExtension == ".mp4"
		if isVideoFile {
			videoFile := models.FileTreeItem{
				Id:   uuid.New().String(),
				Path: getFolderPath(currentItemPath),
				Name: getFileNameWithoutExtension(itemName),
				Type: "video",
			}

			currentFileItems = append(currentFileItems, videoFile)
			continue
		}

		isSubtitleFile := fileExtension == ".vtt"
		if !isSubtitleFile {
			continue
		}

		mp3RegExp := regexp.MustCompile(".(mp3|wav).vtt$")
		isAssociatedWithMp3File := mp3RegExp.MatchString(itemName)

		if !isAssociatedWithMp3File {
			log.Default().Printf("'%v' is not associated with an (mp3|wav) file\n", itemName)
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

		audioFileId := uuid.New().String()
		newSubtitleFile := models.FileTreeItem{
			Id:                    uuid.New().String(),
			Path:                  getFolderPath(currentItemPath),
			Name:                  getFileNameWithoutExtension(itemName),
			Type:                  "subtitle",
			AssociatedAudioFileId: audioFileId,
		}

		newAssociatedSourceFile := models.FileTreeItem{
			Id:   audioFileId,
			Path: getFolderPath(correspondingSourceFilePath),
			Name: getFileNameWithoutExtension(itemName),
			Type: "audio",
		}

		currentFileItems = append(currentFileItems, newSubtitleFile)
		currentFileItems = append(currentFileItems, newAssociatedSourceFile)
	}

	return currentFileItems
}
