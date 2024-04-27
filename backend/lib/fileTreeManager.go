package lib

import (
	"backend/enums"
	lib "backend/lib/utilities"
	"backend/models"
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

		isSubtitleFile := fileType == enums.SUBTITLE
		if !isSubtitleFile {
			continue
		}

		correspondingSourceFilePath := strings.Replace(currentItemPath, path.Ext(itemName), "", 1)
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
			Path:                  lib.GetFolderPath(lib.GetFolderPathInput{Path: currentItemPath, RootPath: rootPath}),
			Name:                  lib.GetFilenameWithoutExtension(itemName),
			Type:                  enums.SUBTITLE,
			AssociatedAudioFileId: audioFileId,
		}

		newAssociatedSourceFile := models.FileTreeItem{
			Id:   audioFileId,
			Path: lib.GetFolderPath(lib.GetFolderPathInput{Path: correspondingSourceFilePath, RootPath: rootPath}),
			Name: lib.GetFilenameWithoutExtension(itemName),
			Type: enums.AUDIO,
		}

		currentFileItems = append(currentFileItems, newSubtitleFile)
		currentFileItems = append(currentFileItems, newAssociatedSourceFile)
	}

	return currentFileItems
}
