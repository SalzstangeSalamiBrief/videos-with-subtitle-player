package handlers

import (
	"backend/internal/problemDetailsErrors"
	"backend/pkg/constants"
	"backend/pkg/repositories"
	"backend/pkg/utilities"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type DiscreteFileByIdHandlerConfig struct {
	RootPath           string
	FileTreeRepository *repositories.FileTreeRepository
}

func CreateDiscreteFileByIdHandler(configuration DiscreteFileByIdHandlerConfig) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fileIdString := strings.TrimPrefix(r.URL.Path, "/api/file/discrete/")
		if fileIdString == "" {
			log.Default().Println(fmt.Sprintf("[DiscreteFileByIdHandler]: Error while opening the file with id '%v'\n", fileIdString))
			problemDetailsErrors.NewBadRequestProblemDetails(fmt.Sprintf("The parameter 'fileId' is empty but required. Please provide an id\n")).SendErrorResponse(w)
			return
		}

		discreteFileInTree, getFileTreeItemsError := configuration.FileTreeRepository.GetFileByFileId(fileIdString)
		if getFileTreeItemsError != nil {
			log.Default().Println(getFileTreeItemsError.Error())
			problemDetailsErrors.NewInternalServerErrorProblemDetails(fmt.Sprintf("Could not get file with id='%v'", fileIdString)).SendErrorResponse(w)
			return
		}

		isExtensionSupported := utilities.IsFileExtensionAllowed(discreteFileInTree, constants.AllowedDiscreteFileExtensions...)
		if !isExtensionSupported {
			log.Default().Println(fmt.Sprintf("File with id='%v' has an unsupported extension", fileIdString))
			problemDetailsErrors.NewInternalServerErrorProblemDetails(fmt.Sprintf("Could not get file with id='%v'", fileIdString)).SendErrorResponse(w)
			return
		}

		filePathOnHardDisk := path.Join(configuration.RootPath, discreteFileInTree.Path)
		fileBytes, err := os.ReadFile(filePathOnHardDisk)
		if err != nil {
			log.Default().Println(fmt.Sprintf("[DiscreteFileByIdHandler]: Error while opening the file with id '%v'\n", err.Error()))
			problemDetailsErrors.NewInternalServerErrorProblemDetails(fmt.Sprintf("Could not find the file with id '%v': %v\n", fileIdString, err)).SendErrorResponse(w)
			return
		}

		w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\"", discreteFileInTree.Name))
		mimeType := utilities.GetContentTypeHeaderMimeType(discreteFileInTree)
		w.Header().Add("Content-Type", mimeType)
		w.Write(fileBytes)
	}

}
