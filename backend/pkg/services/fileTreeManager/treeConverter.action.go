package fileTreeManager

import (
	"backend/pkg/enums/fileType"
	"backend/pkg/models"
	"backend/pkg/services/imageConverter/constants"
	imageConverterUtilities "backend/pkg/services/imageConverter/utilities"
	"backend/pkg/utilities"
	"path/filepath"
	"regexp"

	"github.com/google/uuid"
)

// TODO TEST
func (fileTreeManager *FileTreeManager) convertFileNodesToTree() models.FolderNode {

	rootFileHierarchy := models.FolderNode{
		FolderId: uuid.New().String(),
		//TODO FILL FIELDS
		ChildFolders: []models.FolderNode{},
	}

	files := fileTreeManager.GetFiles()
	for _, file := range files {
		pathParts := utilities.GetPartsOfPath(file)
		buildSubFileTree(&rootFileHierarchy, pathParts)
	}

	for _, file := range files {
		pathParts := utilities.GetPartsOfPath(file)
		getThumbnailOfTree(&rootFileHierarchy, file, pathParts)
	}

	for _, file := range files {
		pathParts := utilities.GetPartsOfPath(file)
		addFileToTree(&rootFileHierarchy, file, pathParts)
	}

	return rootFileHierarchy
}

func buildSubFileTree(parentTree *models.FolderNode, pathPartsWithoutFileExtension []string) {
	remainingPathParts := pathPartsWithoutFileExtension
	currentNode := parentTree

	for i := 0; i < len(remainingPathParts); i += 1 {
		currentPathPart := remainingPathParts[i]
		indexOfMatchingChild := findChildIndexInChildrenOfFileTree(currentNode, currentPathPart)

		if indexOfMatchingChild >= 0 {
			currentNode = &currentNode.ChildFolders[indexOfMatchingChild]
			continue
		}
		// TODO THIS CREATES A NEW UUID FOR EACH FOLDER
		child := models.FolderNode{
			FolderId:     uuid.New().String(),
			Name:         currentPathPart,
			Path:         filepath.Join(pathPartsWithoutFileExtension[0 : i+1]...),
			ChildFolders: []models.FolderNode{},
			Files:        []models.FileNode{},
		}
		currentNode.ChildFolders = append(currentNode.ChildFolders, child)
		currentNode = &child
	}
}

func getThumbnailOfTree(rootFileTree *models.FolderNode, file models.FileNode, pathPartsWithFileExtension []string) {
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

func addFileToTree(rootFileTree *models.FolderNode, file models.FileNode, pathPartsWithFileExtension []string) {
	currentNode := getNodeAssociatedWithFileInTree(rootFileTree, pathPartsWithFileExtension)
	currentNode.Files = append(currentNode.Files, file)
}

func getNodeAssociatedWithFileInTree(rootFileTree *models.FolderNode, pathPartsWithFileExtension []string) *models.FolderNode {
	var currentPathPart string
	remainingPathParts := pathPartsWithFileExtension
	currentNode := rootFileTree

	for i := 0; i < len(pathPartsWithFileExtension); i += 1 {
		currentPathPart = remainingPathParts[i]
		indexOfMatchingChild := findChildIndexInChildrenOfFileTree(currentNode, currentPathPart)
		if indexOfMatchingChild >= 0 {
			currentNode = &currentNode.ChildFolders[indexOfMatchingChild]
		}
	}

	return currentNode
}

func findChildIndexInChildrenOfFileTree(node *models.FolderNode, name string) int {
	for i, child := range node.ChildFolders {
		if child.Name == name {
			return i
		}
	}

	return -1
}
