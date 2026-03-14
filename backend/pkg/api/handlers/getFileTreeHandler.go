package handlers

import (
	"backend/internal/problemDetailsErrors"
	"backend/pkg/models"
	"backend/pkg/repositories"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var fileTree models.FolderNodeDto

type FileTreeHandlerConfiguration struct {
	FolderNodeRepository *repositories.FolderNodeRepository
}

func CreateGetFileTreeHandler(configuration FileTreeHandlerConfiguration) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if fileTree.Id == "" {
			folderNodes, getFileNodes := configuration.FolderNodeRepository.GetFolders()
			if getFileNodes != nil {
				log.Default().Println("Could not get file tree from database")
				problemDetailsErrors.NewInternalServerErrorProblemDetails("File tree ist not initialized").SendErrorResponse(w)
				return
			}

			fileTree = models.NodesToSingleTree(folderNodes)
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
