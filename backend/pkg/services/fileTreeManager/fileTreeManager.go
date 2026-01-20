package fileTreeManager

import (
	"backend/pkg/enums"
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

type SubFileTree []models.FileTreeItem

type FileTreeManager struct {
	fileTreeItems []models.FileTreeItem
	rootPath      string
}

func NewFileTreeManager(rootPath string) *FileTreeManager {
	return &FileTreeManager{
		rootPath:      rootPath,
		fileTreeItems: []models.FileTreeItem{},
	}
}

func (fileTreeManager *FileTreeManager) GetTree() []models.FileTreeItem {
	return fileTreeManager.fileTreeItems
}

func (fileTreeManager *FileTreeManager) InitializeTree() *FileTreeManager {
	log.Default().Println("Start file tree initialization")
	fullTree := fileTreeManager.getSubTree(fileTreeManager.rootPath)
	fileTreeManager.fileTreeItems = fullTree
	log.Default().Println("Finish file tree initialization")
	return fileTreeManager
}

func (fileTreeManager *FileTreeManager) getSubTree(parentPath string) []models.FileTreeItem {
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

		fileType := commonUtilities.GetFileType(itemName)
		if fileType == enums.UNKNOWN {
			continue
		}

		newFileItem := models.FileTreeItem{
			FileId: uuid.New().String(),
			Path:   utilities.GetFolderPath(utilities.GetFolderPathInput{Path: currentItemPath, RootPath: fileTreeManager.rootPath}),
			Name:   utilities.GetFilenameWithoutExtension(itemName),
			Type:   fileType,
		}

		if fileType == enums.IMAGE {
			currentFileItems.handleImageFile(fileTreeManager.rootPath, newFileItem)
			continue
		}

		if fileType == enums.VIDEO {
			currentFileItems.handleVideoFile(newFileItem)
			continue
		}

		if fileType == enums.AUDIO {
			currentFileItems.handleAudioFile(fileTreeManager.rootPath, newFileItem, currentItemPath, itemName)
		}
	}

	return currentFileItems
}

func (subTree *SubFileTree) handleDirectory(fileTreeManager FileTreeManager, directoryPath string, directoryName string) {
	log.Default().Printf("'%v' is a directory\n", directoryName)
	newDirectoryItems := fileTreeManager.getSubTree(directoryPath)
	*subTree = append(*subTree, newDirectoryItems...)
}

func (subTree *SubFileTree) handleVideoFile(videoFile models.FileTreeItem) {
	*subTree = append(*subTree, videoFile)
}

func (subTree *SubFileTree) handleAudioFile(rootPath string, audioFile models.FileTreeItem, currentItemPath string, itemName string) {
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

	subtitleFile := models.FileTreeItem{
		FileId: uuid.New().String(),
		Path:   utilities.GetFolderPath(utilities.GetFolderPathInput{Path: possibleSubtitleFileName, RootPath: rootPath}),
		// TODO NAME INCLUDES THE WHOLE PATH
		Name:                  utilities.GetFilenameWithoutExtension(possibleSubtitleFileName),
		Type:                  enums.SUBTITLE,
		AssociatedAudioFileId: &audioFile.FileId,
	}

	*subTree = append(*subTree, subtitleFile)
}

func (subTree *SubFileTree) handleImageFile(rootPath string, imageFile models.FileTreeItem) {
	isLowQualityImage := imageConverterUtilities.IsLowQualityImage(imageFile.Path)
	if isLowQualityImage {
		return
	}

	doesLowQualityImageExist := imageConverterUtilities.DoesLowQualityImageExist(rootPath, imageFile.Path)
	if doesLowQualityImageExist {
		lowQualityImagePath := imageConverterUtilities.GetLowQualityImagePath(imageFile.Path)

		resizeImageFileItem := models.FileTreeItem{
			FileId: uuid.New().String(),
			Path:   lowQualityImagePath,
			Name:   utilities.GetFilenameWithoutExtension(lowQualityImagePath),
			Type:   enums.IMAGE,
		}
		imageFile.LowQualityImageId = &resizeImageFileItem.FileId
		*subTree = append(*subTree, resizeImageFileItem)
	}

	*subTree = append(*subTree, imageFile)
}
