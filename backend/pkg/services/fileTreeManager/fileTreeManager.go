package fileTreeManager

import (
	"backend/internal/config"
	"backend/pkg/enums"
	"backend/pkg/models"
	"backend/pkg/services/ImageQualityReducer"
	"backend/pkg/services/fileTreeManager/utilities"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type SubFileTree []models.FileTreeItem

var FileTreeItems []models.FileTreeItem

func InitializeFileTree() {
	fullTree := getFullTree(config.AppConfiguration.RootPath)
	FileTreeItems = fullTree
	log.Default().Println("Finish file tree initialization")
}

func getFullTree(parentPath string) []models.FileTreeItem {
	content, err := os.ReadDir(parentPath)
	if err != nil {
		log.Fatal(err)
	}

	currentFileItems := make(SubFileTree, 0)
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

		if fileType == enums.IMAGE {
			currentFileItems.HandleImageFile(newFileItem, currentItemPath)
			continue
		}

		if fileType == enums.VIDEO {
			currentFileItems.HandleVideoFile(newFileItem)
			continue
		}

		if fileType == enums.AUDIO {
			currentFileItems.HandleAudioFile(newFileItem, currentItemPath, itemName)
		}
	}

	return currentFileItems
}

func (input *SubFileTree) HandleVideoFile(videoFile models.FileTreeItem) {
	*input = append(*input, videoFile)
}

func (input *SubFileTree) HandleAudioFile(audioFile models.FileTreeItem, currentItemPath string, itemName string) {
	*input = append(*input, audioFile)

	possibleSubtitleFileName := strings.Replace(currentItemPath, path.Ext(itemName), fmt.Sprintf("%v.vtt", path.Ext(itemName)), 1)
	_, err := os.Stat(possibleSubtitleFileName)
	if err != nil {
		log.Default().Printf("error while checking if a matching subttile file exists. Sourcefile '%v'; Error: '%v'\n", itemName, err.Error())
		return
	}

	isNotAssociatedWithSubtitleFile := os.IsNotExist(err)
	if isNotAssociatedWithSubtitleFile {
		log.Default().Printf("No matching subtitle file for audio file '%v' exists\n", itemName)
		return
	}

	subtitleFile := models.FileTreeItem{
		Id:   uuid.New().String(),
		Path: utilities.GetFolderPath(utilities.GetFolderPathInput{Path: possibleSubtitleFileName, RootPath: config.AppConfiguration.RootPath}),
		// TODO NAME INCLUDES THE WHOLE PATH
		Name:                  utilities.GetFilenameWithoutExtension(possibleSubtitleFileName),
		Type:                  enums.SUBTITLE,
		AssociatedAudioFileId: audioFile.Id,
	}

	*input = append(*input, subtitleFile)
}

func (input *SubFileTree) HandleImageFile(imageFile models.FileTreeItem, currentItemPath string) {
	isLowQualityImage := ImageQualityReducer.IsLowQualityFileName(currentItemPath)
	if isLowQualityImage {
		log.Default().Printf("'%v' is already a low quality image\n", imageFile.Name)
		return
	}

	lowQualityImagePath, err := ImageQualityReducer.ReduceImageQuality(currentItemPath)
	if err != nil {
		log.Default().Printf("Error reducing the quality of the image '%v': %v\n", imageFile.Path, err.Error())
		return
	}

	resizeImageFileItem := models.FileTreeItem{
		Id:   uuid.New().String(),
		Path: utilities.GetFolderPath(utilities.GetFolderPathInput{Path: lowQualityImagePath, RootPath: config.AppConfiguration.RootPath}),
		Name: utilities.GetFilenameWithoutExtension(lowQualityImagePath),
		Type: enums.IMAGE,
	}

	imageFile.LowQualityImageId = resizeImageFileItem.Id
	*input = append(*input, imageFile)
	*input = append(*input, resizeImageFileItem)
}
