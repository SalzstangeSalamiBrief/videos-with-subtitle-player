package usecases

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"videos-with-subtitle-player/models"
	"videos-with-subtitle-player/router"
	directoryTree "videos-with-subtitle-player/services/directoryTree"
)

var GetFileTreeUseCaseRoute = router.Route{
	Path:    "/file-tree",
	Handler: getFileTreeUseCase,
	Method:  http.MethodGet,
}

func getFileTreeUseCase(w http.ResponseWriter, r *http.Request, quit chan<- bool) {
	fileTree := directoryTree.FileTreeItems
	encodedBytes, err := json.Marshal(fileTree)
	if err != nil {
		router.ErrorHandler(w, fmt.Sprintf("Could not marshal file tree: %v", err.Error()), http.StatusInternalServerError)
		quit <- true
		return
	}

	w.Write(encodedBytes)
	quit <- true
}

func getFileTreeDto(flatFileTree []models.FileTreeItem) {
	rootFileHierarchy := models.FileTreeDto{
		Children: map[string]models.FileTreeDto{},
	}

	for _, file := range flatFileTree {
		pathParts := strings.Split(file.Path, "/")[1:]                // first element is empty, so skip it
		pathPartsWithoutFileExtension := pathParts[:len(pathParts)-1] // remove file extension
		hierarchyItem := buildFileHierarchyItemOfPathParts(pathPartsWithoutFileExtension[1:])
		rootFileHierarchy.Children[pathParts[0]] = hierarchyItem
	}

	for _, file := range flatFileTree {
		pathParts := strings.Split(file.Path, "/")[1:]                // first element is empty, so skip it
		pathPartsWithoutFileExtension := pathParts[:len(pathParts)-1] // remove file extension
		AddFileToHierarchy(&rootFileHierarchy, file, pathPartsWithoutFileExtension)
	}
	fmt.Println(rootFileHierarchy)
}

func buildFileHierarchyItemOfPathParts(pathParts []string) models.FileTreeDto {
	item := models.FileTreeDto{
		Children:   map[string]models.FileTreeDto{},
		AudioFiles: map[string][]models.FileItem{},
	}

	if len(pathParts) == 0 {
		return item
	}

	currentPathPart := pathParts[0]
	remainingPathParts := pathParts[1:]

	_, ok := item.Children[currentPathPart]
	if !ok {
		item.Children[currentPathPart] = models.FileTreeDto{
			Children:   map[string]models.FileTreeDto{},
			AudioFiles: map[string][]models.FileItem{},
		}
	}

	// Update the children by making a recursive call with the correct remainingPathParts
	item.Children[currentPathPart] = buildFileHierarchyItemOfPathParts(remainingPathParts)
	return item
}

func AddFileToHierarchy(currentHierarchy *models.FileTreeDto, file models.FileTreeItem, remainingPathParts []string) {

	if len(remainingPathParts) == 1 {
		matchingHierarchy := (*currentHierarchy).Children[remainingPathParts[0]]

		if matchingHierarchy.AudioFiles == nil {
			matchingHierarchy.AudioFiles = map[string][]models.FileItem{}
		}

		_, ok := matchingHierarchy.AudioFiles[file.Name]
		if ok == false {
			matchingHierarchy.AudioFiles[file.Name] = []models.FileItem{}
		}

		fileDto := models.FileItem{
			Id:   file.Id,
			Name: file.Name,
		}

		matchingHierarchy.AudioFiles[file.Name] = append(matchingHierarchy.AudioFiles[file.Name], fileDto)
		return
	}

	currentPathPart, remainingPath := remainingPathParts[0], remainingPathParts[1:]
	// Ensure the current path part has a corresponding child hierarchy
	if currentHierarchy.Children == nil {
		currentHierarchy.Children = map[string]models.FileTreeDto{}
	}

	nextHierarchy, ok := currentHierarchy.Children[currentPathPart]
	if !ok {
		nextHierarchy = models.FileTreeDto{
			Children: map[string]models.FileTreeDto{},
		}
		currentHierarchy.Children[currentPathPart] = nextHierarchy
	}

	// Recursively add the file to the next hierarchy
	AddFileToHierarchy(&nextHierarchy, file, remainingPath)
}
