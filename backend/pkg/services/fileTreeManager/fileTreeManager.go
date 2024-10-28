package fileTreeManager

import (
	"backend/internal/config"
	"backend/pkg/enums"
	"backend/pkg/models"
	"backend/pkg/services/fileTreeManager/utilities"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var FileTreeItems []models.FileTreeItem

func InitializeFileTree() {
	fullTree := getFullTree(config.AppConfiguration.RootPath)
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
		// TODO REFACTOR INTO SWITCH WITH FUNCTION CALLS
		if isDirectory {
			log.Default().Printf("'%v' is a directory\n", itemName)
			newDirectoryItems := getFullTree(currentItemPath)
			currentFileItems = append(currentFileItems, newDirectoryItems...)
			continue
		}

		fileType := utilities.GetFileType(itemName)
		if fileType == enums.UNKNOWN {
			continue
		}

		newFileItem := models.FileTreeItem{
			Id:   uuid.New().String(),
			Path: utilities.GetFolderPath(utilities.GetFolderPathInput{Path: currentItemPath, RootPath: config.AppConfiguration.RootPath}),
			Name: utilities.GetFilenameWithoutExtension(itemName),
			Type: fileType,
		}

		if fileType == enums.VIDEO || fileType == enums.IMAGE {
			currentFileItems = append(currentFileItems, newFileItem)
			continue
		}

		if fileType == enums.AUDIO {
			currentFileItems = append(currentFileItems, newFileItem)

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
				Path: utilities.GetFolderPath(utilities.GetFolderPathInput{Path: possibleSubtitleFileName, RootPath: config.AppConfiguration.RootPath}),
				// TODO NAME INCLUDES THE WHOLE PATH
				Name:                  utilities.GetFilenameWithoutExtension(possibleSubtitleFileName),
				Type:                  enums.SUBTITLE,
				AssociatedAudioFileId: newFileItem.Id,
			}
			currentFileItems = append(currentFileItems, subtitleFile)

		}
	}

	return currentFileItems
}
