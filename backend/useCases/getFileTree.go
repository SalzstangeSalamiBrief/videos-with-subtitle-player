package usecases

import (
	"backend/lib"
	"backend/models"
	"backend/models/enums"
	"backend/router"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/google/uuid"
)

var GetFileTreeUseCaseRoute = router.Route{
	Path:    "/file-tree",
	Handler: getFileTreeUseCase,
	Method:  http.MethodGet,
}

func getFileTreeUseCase(w http.ResponseWriter, r *http.Request) {
	fileTree := GetFileTreeDto(lib.FileTreeItems)
	encodedBytes, err := json.Marshal(fileTree.Children)
	if err != nil {
		router.ErrorHandler(w, fmt.Sprintf("Could not marshal file tree: %v", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Write(encodedBytes)
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
				Id:       uuid.New().String(),
				Name:     currentPathPart,
				Children: []models.FileTreeDto{},
				Files:    []models.FileDto{},
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
	currentNode := getNodeAssociatedWithFileInTree(rootFileTree, pathPartsWithFileExtension)

	fileItem := models.FileDto{
		Id:                    file.Id,
		Name:                  file.Name,
		Type:                  file.Type,
		AssociatedAudioFileId: file.AssociatedAudioFileId,
	}

	currentNode.Files = append(currentNode.Files, fileItem)
}

func getNodeAssociatedWithFileInTree(rootFileTree *models.FileTreeDto, pathPartsWithFileExtension []string) *models.FileTreeDto {
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

	return currentNode
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
	for i, child := range node.Files {
		if child.Name == name {
			return i
		}
	}
	return -1
}

func getFileType(fileName string) (enums.FileType, error) {
	extension := path.Ext(fileName)
	switch extension {
	case ".mp3":
		return enums.AUDIO, nil
	case ".mp4":
		return enums.VIDEO, nil
	case ".wav":
		return enums.AUDIO, nil
	case ".vtt":
		return enums.SUBTITLE, nil
	default:
		return enums.UNKNOWN, fmt.Errorf("unknown file type for file '%v'", fileName)
	}
}
