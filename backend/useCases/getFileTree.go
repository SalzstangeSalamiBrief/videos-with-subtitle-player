package usecases

import (
	"backend/enums"
	"backend/lib"
	"backend/models"
	"backend/router"
	usecases "backend/useCases/utilities"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"regexp"
)

var GetFileTreeUseCaseRoute = router.Route{
	Path:    "/file-tree",
	Handler: getFileTreeUseCase,
	Method:  http.MethodGet,
}

var fileTree models.FileTreeDto

func getFileTreeUseCase(w http.ResponseWriter, r *http.Request) {
	if fileTree.Id == "" {
		fileTree = GetFileTreeDto(lib.FileTreeItems)
	}

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
		pathParts := usecases.GetPartsOfPath(file)
		buildSubFileTree(&rootFileHierarchy, pathParts)
	}

	for _, file := range filesArray {
		pathParts := usecases.GetPartsOfPath(file)
		getThumbnailOfTree(&rootFileHierarchy, file, pathParts)
	}

	for _, file := range filesArray {
		pathParts := usecases.GetPartsOfPath(file)
		addFileToTree(&rootFileHierarchy, file, pathParts)
	}

	return rootFileHierarchy
}

func buildSubFileTree(parentTree *models.FileTreeDto, pathPartsWithoutFileExtension []string) {
	remainingPathParts := pathPartsWithoutFileExtension
	currentNode := parentTree

	for i := 0; i < len(remainingPathParts); i += 1 {
		currentPathPart := remainingPathParts[i]
		indexOfMatchingChild := findChildIndexInChildrenOfFileTree(currentNode, currentPathPart)

		if indexOfMatchingChild >= 0 {
			currentNode = &currentNode.Children[indexOfMatchingChild]
			continue
		}

		child := models.FileTreeDto{
			Id:       currentNode.Id,
			Name:     currentPathPart,
			Children: []models.FileTreeDto{},
			Files:    []models.FileDto{},
		}
		currentNode.Children = append(currentNode.Children, child)
		currentNode = &child
	}

	//for isGettingMatchingItemInHierarchy {
	//	indexOfMatchingChild := findChildIndexInChildrenOfFileTree(currentNode, remainingPathParts[0])
	//
	//	hasMatchingChild := indexOfMatchingChild >= 0
	//	if hasMatchingChild {
	//		currentNode = &currentNode.Children[indexOfMatchingChild]
	//	} else {
	//		child := models.FileTreeDto{
	//			Id:       currentNode.Id,
	//			Name:     currentPathPart,
	//			Children: []models.FileTreeDto{},
	//			Files:    []models.FileDto{},
	//		}
	//		currentNode.Children = append(currentNode.Children, child)
	//		currentNode = &child
	//	}
	//
	//	remainingPathParts = remainingPathParts[1:]
	//	isGettingMatchingItemInHierarchy = len(remainingPathParts) > 0
	//}
}

func getThumbnailOfTree(rootFileTree *models.FileTreeDto, file models.FileTreeItem, pathPartsWithFileExtension []string) {
	if file.Type != enums.IMAGE {
		return
	}

	currentNode := getNodeAssociatedWithFileInTree(rootFileTree, pathPartsWithFileExtension)
	if currentNode.ThumbnailId != "" {
		return
	}

	isThumbnailImage := regexp.MustCompile(`(?i)Thumbnail`).MatchString(file.Name)
	if !isThumbnailImage {
		return
	}

	currentNode.ThumbnailId = file.Id
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
