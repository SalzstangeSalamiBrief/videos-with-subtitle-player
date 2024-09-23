package lib

import (
	"backend/enums"
	lib "backend/lib/utilities"
	"backend/models"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"path"
	"path/filepath"
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
		currentItemPath := filepath.Join(parentPath, itemName)
		isDirectory := item.IsDir()

		if isDirectory {
			log.Default().Printf("'%v' is a directory\n", itemName)
			newDirectoryItems := getFullTree(currentItemPath)
			currentFileItems = append(currentFileItems, newDirectoryItems...)
			continue
		}

		fileType := lib.GetFileType(itemName)
		if fileType == enums.UNKNOWN {
			continue
		}

		isVideoFile := fileType == enums.VIDEO
		if isVideoFile {
			videoFile := models.FileTreeItem{
				Id:   uuid.New().String(),
				Path: lib.GetFolderPath(lib.GetFolderPathInput{Path: currentItemPath, RootPath: rootPath}),
				Name: lib.GetFilenameWithoutExtension(itemName),
				Type: fileType,
			}

			currentFileItems = append(currentFileItems, videoFile)
			continue
		}

		isAudioFile := fileType == enums.AUDIO
		if isAudioFile {
			audioFile := models.FileTreeItem{
				Id:   uuid.New().String(),
				Path: lib.GetFolderPath(lib.GetFolderPathInput{Path: currentItemPath, RootPath: rootPath}),
				Name: lib.GetFilenameWithoutExtension(itemName),
				Type: fileType,
			}
			currentFileItems = append(currentFileItems, audioFile)

			possibleSubtitleFileName := strings.Replace(currentItemPath, path.Ext(itemName), fmt.Sprintf("%v.vtt", path.Ext(itemName)), 1)
			_, err := os.Stat(possibleSubtitleFileName)
			if err != nil {
				log.Default().Printf("error while checking if a matching subttile file exists. Sourcefile '%v'; Error: '%v'\n", itemName, err.Error())
				continue
			}

			isNotAssociatedWithSubtitleFile := os.IsNotExist(err)
			if isNotAssociatedWithSubtitleFile {
				log.Default().Printf("No matching subtitle file for audio file '%v' exists\n", itemName)
				continue
			}
			subtitleFile := models.FileTreeItem{
				Id:   uuid.New().String(),
				Path: lib.GetFolderPath(lib.GetFolderPathInput{Path: possibleSubtitleFileName, RootPath: rootPath}),
				// TODO NAME INCLUDES THE WHOLE PATH
				Name:                  lib.GetFilenameWithoutExtension(possibleSubtitleFileName),
				Type:                  enums.SUBTITLE,
				AssociatedAudioFileId: audioFile.Id,
			}
			currentFileItems = append(currentFileItems, subtitleFile)

		}

	}

	return currentFileItems
}
