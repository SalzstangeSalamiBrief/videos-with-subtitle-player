package fileTreeManager

import (
	"backend/pkg/enums/fileType"
	"backend/pkg/models"
	"backend/pkg/services/fileTreeManager/utilities"
	imageConverterUtilities "backend/pkg/services/imageConverter/utilities"
	commonUtilities "backend/pkg/utilities"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func (fileTreeManager *FileTreeManager) scanForFilesInPath(parentPath string) []models.FileNode {
	content, err := os.ReadDir(parentPath)
	if err != nil {
		log.Fatal(err)
	}

	currentFileItems := make(SubFileTree, 0)
	for _, item := range content {
		itemName := item.Name()
		currentItemPath := filepath.Join(parentPath, itemName)
		isDirectory := item.IsDir()
		if isDirectory {
			currentFileItems.handleDirectory(*fileTreeManager, currentItemPath, itemName)
			continue
		}

		fileTypeOfItem := commonUtilities.GetFileType(itemName)
		if fileTypeOfItem == fileType.UNKNOWN {
			continue
		}

		newFileItem := models.FileNode{
			FileId: uuid.New().String(),
			Path:   utilities.GetFolderPath(utilities.GetFolderPathInput{Path: currentItemPath, RootPath: fileTreeManager.rootPath}),
			Name:   utilities.GetFilenameWithoutExtension(itemName),
			Type:   fileTypeOfItem,
		}

		if fileTypeOfItem == fileType.IMAGE {
			currentFileItems.handleImageFile(fileTreeManager.rootPath, newFileItem)
			continue
		}

		if fileTypeOfItem == fileType.VIDEO {
			currentFileItems.handleVideoFile(newFileItem)
			continue
		}

		if fileTypeOfItem == fileType.AUDIO {
			currentFileItems.handleAudioFile(fileTreeManager.rootPath, newFileItem, currentItemPath, itemName)
		}
	}

	return currentFileItems
}

func (subTree *SubFileTree) handleDirectory(fileTreeManager FileTreeManager, directoryPath string, directoryName string) {
	log.Default().Printf("'%v' is a directory\n", directoryName)
	newDirectoryItems := fileTreeManager.scanForFilesInPath(directoryPath)
	*subTree = append(*subTree, newDirectoryItems...)
}

func (subTree *SubFileTree) handleVideoFile(videoFile models.FileNode) {
	*subTree = append(*subTree, videoFile)
}

func (subTree *SubFileTree) handleAudioFile(rootPath string, audioFile models.FileNode, currentItemPath string, itemName string) {
	*subTree = append(*subTree, audioFile)

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

	subtitleFile := models.FileNode{
		FileId: uuid.New().String(),
		Path:   utilities.GetFolderPath(utilities.GetFolderPathInput{Path: possibleSubtitleFileName, RootPath: rootPath}),
		// TODO NAME INCLUDES THE WHOLE PATH
		Name:                  utilities.GetFilenameWithoutExtension(possibleSubtitleFileName),
		Type:                  fileType.SUBTITLE,
		AssociatedAudioFileId: &audioFile.FileId,
	}

	*subTree = append(*subTree, subtitleFile)
}

func (subTree *SubFileTree) handleImageFile(rootPath string, imageFile models.FileNode) {
	isLowQualityImage := imageConverterUtilities.IsLowQualityImagePath(imageFile.Path)
	if isLowQualityImage {
		return
	}

	doesLowQualityImageExist := imageConverterUtilities.DoesLowQualityImageExist(rootPath, imageFile.Path)
	if doesLowQualityImageExist {
		lowQualityImagePath := imageConverterUtilities.GetLowQualityImagePath(imageFile.Path)

		resizeImageFileItem := models.FileNode{
			FileId: uuid.New().String(),
			Path:   lowQualityImagePath,
			Name:   utilities.GetFilenameWithoutExtension(lowQualityImagePath),
			Type:   fileType.IMAGE,
		}
		imageFile.LowQualityImageId = &resizeImageFileItem.FileId
		*subTree = append(*subTree, resizeImageFileItem)
	}

	*subTree = append(*subTree, imageFile)
}
