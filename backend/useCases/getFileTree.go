package usecases

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"
	"videos-with-subtitle-player/models"
	"videos-with-subtitle-player/router"
	directoryTree "videos-with-subtitle-player/services/directoryTree"

	"github.com/google/uuid"
)

var GetFileTreeUseCaseRoute = router.Route{
	Path:    "/file-tree",
	Handler: getFileTreeUseCase,
	Method:  http.MethodGet,
}

func getFileTreeUseCase(w http.ResponseWriter, r *http.Request, quit chan<- bool) {
	fileTree := GetFileTreeDto(directoryTree.FileTreeItems)
	subTrees := fileTree.Children
	encodedBytes, err := json.Marshal(subTrees)
	if err != nil {
		router.ErrorHandler(w, fmt.Sprintf("Could not marshal file tree: %v", err.Error()), http.StatusInternalServerError)
		quit <- true
		return
	}

	w.Write(encodedBytes)
	quit <- true
}

func GetFileTreeDto(filesArray []models.FileTreeItem) models.FileTreeDto {
	rootFileHierarchy := models.FileTreeDto{
		Id:       uuid.New().String(),
		Children: []models.FileTreeDto{},
	}

	for _, file := range filesArray {
		pathParts := strings.Split(file.Path, "/")[1:]                // first element is empty, so skip it
		pathPartsWithoutFileExtension := pathParts[:len(pathParts)-1] // remove file extension
		buildSubFileTree(&rootFileHierarchy, pathPartsWithoutFileExtension)
	}

	for _, file := range filesArray {
		pathParts := strings.Split(file.Path, "/")[1:] // first element is empty, so skip it
		addFileToTree(&rootFileHierarchy, file, pathParts)
	}

	return rootFileHierarchy
}

func buildSubFileTree(parentTree *models.FileTreeDto, pathPartsWithoutFileExtension []string) {
	var currentPathPart string
	remainingPathParts := pathPartsWithoutFileExtension
	currentNode := parentTree

	isGettingMatchingItemInHierarchy := len(remainingPathParts[0]) > 0
	for isGettingMatchingItemInHierarchy {
		currentPathPart = remainingPathParts[0]
		indexOfMatchingChild := findChildIndexInChildrenOfFileTree(currentNode, currentPathPart)

		hasMatchingChild := indexOfMatchingChild >= 0
		if hasMatchingChild {
			currentNode = &currentNode.Children[indexOfMatchingChild]
		} else {
			child := models.FileTreeDto{
				Id:         uuid.New().String(),
				Name:       currentPathPart,
				Children:   []models.FileTreeDto{},
				AudioFiles: []models.AudioFileDto{},
			}
			currentNode.Children = append(currentNode.Children, child)
			currentNode = &child
		}

		remainingPathParts = remainingPathParts[1:]
		if len(remainingPathParts) == 0 {
			isGettingMatchingItemInHierarchy = false
		}
	}
}

func addFileToTree(rootFileTree *models.FileTreeDto, file models.FileTreeItem, pathPartsWithFileExtension []string) {
	var currentPathPart string
	remainingPathParts := pathPartsWithFileExtension
	currentNode := rootFileTree

	for len(remainingPathParts) > 0 {
		currentPathPart = remainingPathParts[0]
		indexOfMatchingChild := findChildIndexInChildrenOfFileTree(currentNode, currentPathPart)
		hasMatchingChild := indexOfMatchingChild >= 0 // cannot have match on the last element because the file name with extension is not included in the file tree

		if hasMatchingChild {
			currentNode = &currentNode.Children[indexOfMatchingChild]
		}

		remainingPathParts = remainingPathParts[1:]
	}

	fileItem := models.FileItemDto{
		Id:   file.Id,
		Name: file.Name,
	}

	indexOfAudioContainer := findChildIndexInAudioFilesOfFileTree(currentNode, fileItem.Name)
	if indexOfAudioContainer < 0 {
		currentNode.AudioFiles = append(currentNode.AudioFiles,
			models.AudioFileDto{
				Name:         fileItem.Name,
				AudioFile:    models.FileItemDto{},
				SubtitleFile: models.FileItemDto{},
			})
		indexOfAudioContainer = len(currentNode.AudioFiles) - 1
	}

	extension := path.Ext(pathPartsWithFileExtension[len(pathPartsWithFileExtension)-1])
	isSubtitleFile := extension == ".vtt"

	if isSubtitleFile {
		currentNode.AudioFiles[indexOfAudioContainer].SubtitleFile = fileItem
	} else {
		currentNode.AudioFiles[indexOfAudioContainer].AudioFile = fileItem
	}
}

func findChildIndexInChildrenOfFileTree(node *models.FileTreeDto, name string) int {
	for i, child := range node.Children {
		if child.Name == name {
			return i
		}
	}
	return -1
}

func findChildIndexInAudioFilesOfFileTree(node *models.FileTreeDto, name string) int {
	for i, child := range node.AudioFiles {
		if child.Name == name {
			return i
		}
	}
	return -1
}
