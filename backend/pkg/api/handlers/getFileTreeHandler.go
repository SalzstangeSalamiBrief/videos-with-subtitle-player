package handlers

import (
	"backend/internal/problemDetailsErrors"
	"backend/pkg/enums/fileType"
	"backend/pkg/models"
	"backend/pkg/repositories"
	"backend/pkg/services/imageConverter/constants"
	imageConverterUtilities "backend/pkg/services/imageConverter/utilities"
	"backend/pkg/utilities"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"regexp"

	"github.com/google/uuid"
)

var fileTree models.FolderNodeDto

type FileTreeHandlerConfiguration struct {
	FileTreeRepository *repositories.FileTreeRepository
}

func CreateGetFileTreeHandler(configuration FileTreeHandlerConfiguration) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if fileTree.Id == "" {
			fileTreeItems, getFileTreeItemsError := configuration.FileTreeRepository.GetFileTree()
			if getFileTreeItemsError != nil {
				log.Default().Println("Could not get file tree from database")
				problemDetailsErrors.NewInternalServerErrorProblemDetails("File tree ist not initialized").SendErrorResponse(w)
				return
			}

			fileTree = getFileTreeDto(fileTreeItems)
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

func getFileTreeDto(filesArray []models.FileNode) models.FolderNodeDto {
	rootFileHierarchy := models.FolderNodeDto{
		Id:       uuid.New().String(),
		Children: []models.FolderNodeDto{},
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

func buildSubFileTree(parentTree *models.FolderNodeDto, pathPartsWithoutFileExtension []string) {
	remainingPathParts := pathPartsWithoutFileExtension
	currentNode := parentTree

	for i := 0; i < len(remainingPathParts); i += 1 {
		currentPathPart := remainingPathParts[i]
		indexOfMatchingChild := findChildIndexInChildrenOfFileTree(currentNode, currentPathPart)

		if indexOfMatchingChild >= 0 {
			currentNode = &currentNode.Children[indexOfMatchingChild]
			continue
		}
		// TODO THIS CREATES A NEW UUID FOR EACH FOLDER
		child := models.FolderNodeDto{
			Id:       uuid.New().String(),
			Name:     currentPathPart,
			Children: []models.FolderNodeDto{},
			Files:    []models.FileNodeDto{},
		}
		currentNode.Children = append(currentNode.Children, child)
		currentNode = &child
	}
}

func getThumbnailOfTree(rootFileTree *models.FolderNodeDto, file models.FileNode, pathPartsWithFileExtension []string) {
	if file.Type != fileType.IMAGE {
		return
	}

	//// TODO ONLY USE WEBP => IMPORT HELPER FROM THE WEBP PACKAGE
	//if filepath.Ext(file.Name) != constants.WebpExtension {
	//	return
	//}

	isLowQualityImage := imageConverterUtilities.IsLowQualityImagePath(file.Path)
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

	currentNode.ThumbnailId = file.FileId
	if file.LowQualityImageId != nil {
		currentNode.LowQualityThumbnailId = *file.LowQualityImageId
	}
}

func addFileToTree(rootFileTree *models.FolderNodeDto, file models.FileNode, pathPartsWithFileExtension []string) {
	currentNode := getNodeAssociatedWithFileInTree(rootFileTree, pathPartsWithFileExtension)

	fileItem := file.ToDto()

	currentNode.Files = append(currentNode.Files, fileItem)
}

func getNodeAssociatedWithFileInTree(rootFileTree *models.FolderNodeDto, pathPartsWithFileExtension []string) *models.FolderNodeDto {
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

func findChildIndexInChildrenOfFileTree(node *models.FolderNodeDto, name string) int {
	for i, child := range node.Children {
		if child.Name == name {
			return i
		}
	}
	return -1
}
