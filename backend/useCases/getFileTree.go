package usecases

import (
	"backend/lib"
	"backend/models"
	"backend/router"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
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
		pathPartsWithoutFileExtension := getPartsOfPath(file)
		buildSubFileTree(&rootFileHierarchy, pathPartsWithoutFileExtension)
	}

	for _, file := range filesArray {
		pathParts := getPartsOfPath(file) // first element is empty, so skip it
		addFileToTree(&rootFileHierarchy, file, pathParts)
	}

	return rootFileHierarchy
}

func getPartsOfPath(file models.FileTreeItem) []string {
	filePath, _ := filepath.Split(file.Path)
	allParts := strings.Split(filePath, string(filepath.Separator))
	var parts []string
	for _, part := range allParts {
		if part == "" {
			continue
		}

		parts = append(parts, part)
	}

	return parts
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
