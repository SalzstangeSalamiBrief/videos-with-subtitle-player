package handlers

import (
	"backend/internal/problemDetailsErrors"
	"backend/pkg/constants"
	"backend/pkg/services/fileTreeManager"
	"backend/pkg/utilities"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type DiscreteFileByIdHandlerConfig struct {
	RootPath string
	*fileTreeManager.FileTreeManager
}

func CreateDiscreteFileByIdHandler(config DiscreteFileByIdHandlerConfig) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fileIdString := strings.TrimPrefix(r.URL.Path, "/api/file/discrete/")
		if fileIdString == "" {
			log.Default().Println(fmt.Sprintf("[DiscreteFileByIdHandler]: Error while opening the file with id '%v'\n", fileIdString))
			problemDetailsErrors.NewBadRequestProblemDetails(fmt.Sprintf("The parameter 'fileId' is empty but required. Please provide an id\n")).SendErrorResponse(w)
			return
		}

		discreteFileInTree := utilities.GetFileByIdAndExtension(config.FileTreeManager.GetTree(), fileIdString, constants.AllowedDiscreteFileExtensions...)
		if discreteFileInTree.FileId == "" {
			log.Default().Println(fmt.Sprintf("[DiscreteFileByIdHandler]: Error while opening the file with id '%v'\n", fileIdString))
			problemDetailsErrors.NewNotFoundProblemDetails(fmt.Sprintf("Could not find file with id '%v'\n", fileIdString)).SendErrorResponse(w)
			return
		}

		filePathOnHardDisk := path.Join(config.RootPath, discreteFileInTree.Path)
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
