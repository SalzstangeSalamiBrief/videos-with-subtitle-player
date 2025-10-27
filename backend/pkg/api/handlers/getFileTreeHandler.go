package handlers

import (
	"backend/internal/problemDetailsErrors"
	"backend/pkg/enums"
	"backend/pkg/models"
	"backend/pkg/services/fileTreeManager"
	"backend/pkg/services/imageConverter/constants"
	utilities2 "backend/pkg/services/imageConverter/utilities"
	"backend/pkg/utilities"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
)

var fileTree models.FileTreeDto

type FileTreeHandlerConfiguration struct {
	FileTreeManager *fileTreeManager.FileTreeManager
}

func CreateGetFileTreeHandler(configuration FileTreeHandlerConfiguration) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if fileTree.Id == "" {
			fileTree = getFileTreeDto(configuration.FileTreeManager.GetTree())
		}

		encodedBytes, err := json.Marshal(fileTree.Children)
		if err != nil {
			log.Default().Println(fmt.Sprintf("Could not marshal file tree: %v", err.Error()))
			problemDetailsErrors.NewInternalServerErrorProblemDetails("Could not marshal file tree").SendErrorResponse(w)

			return
		}

		w.Write(encodedBytes)
	}
}

func getFileTreeDto(filesArray []models.FileTreeItem) models.FileTreeDto {
	rootFileHierarchy := models.FileTreeDto{
		Id:       uuid.New().String(),
		Children: []models.FileTreeDto{},
	}

	for _, file := range filesArray {
		pathParts := utilities.GetPartsOfPath(file)
		buildSubFileTree(&rootFileHierarchy, pathParts)
	}

	for _, file := range filesArray {
		pathParts := utilities.GetPartsOfPath(file)
		getThumbnailOfTree(&rootFileHierarchy, file, pathParts)
	}

	for _, file := range filesArray {
		pathParts := utilities.GetPartsOfPath(file)
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
			Id:       uuid.New().String(),
			Name:     currentPathPart,
			Children: []models.FileTreeDto{},
			Files:    []models.FileDto{},
		}
		currentNode.Children = append(currentNode.Children, child)
		currentNode = &child
	}
}

func getThumbnailOfTree(rootFileTree *models.FileTreeDto, file models.FileTreeItem, pathPartsWithFileExtension []string) {
	if file.Type != enums.IMAGE {
		return
	}

	//// TODO ONLY USE WEBP => IMPORT HELPER FROM THE WEBP PACKAGE
	//if filepath.Ext(file.Name) != constants.WebpExtension {
	//	return
	//}

	isLowQualityImage := utilities2.IsLowQualityImage(file.Path)
	if isLowQualityImage {
		return
	}

	isWebP := filepath.Ext(file.Path) == constants.WebpExtension
	if !isWebP {
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
	currentNode.LowQualityThumbnailId = file.LowQualityImageId
}

func addFileToTree(rootFileTree *models.FileTreeDto, file models.FileTreeItem, pathPartsWithFileExtension []string) {
	currentNode := getNodeAssociatedWithFileInTree(rootFileTree, pathPartsWithFileExtension)

	fileItem := models.FileDto{
		Id:                    file.Id,
		Name:                  file.Name,
		Type:                  file.Type,
		AssociatedAudioFileId: file.AssociatedAudioFileId,
		LowQualityImageId:     file.LowQualityImageId,
	}

	currentNode.Files = append(currentNode.Files, fileItem)
}

func getNodeAssociatedWithFileInTree(rootFileTree *models.FileTreeDto, pathPartsWithFileExtension []string) *models.FileTreeDto {
	var currentPathPart string
	remainingPathParts := pathPartsWithFileExtension
	currentNode := rootFileTree

	for i := 0; i < len(pathPartsWithFileExtension); i += 1 {
		currentPathPart = remainingPathParts[i]
		indexOfMatchingChild := findChildIndexInChildrenOfFileTree(currentNode, currentPathPart)
		if indexOfMatchingChild >= 0 {
			currentNode = &currentNode.Children[indexOfMatchingChild]
		}
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
