package handlers

import (
	"backend/pkg/constants"
	"backend/pkg/services/fileTreeManager"
	"backend/pkg/utilities"
	"fmt"
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
		discreteFileInTree := utilities.GetFileByIdAndExtension(config.FileTreeManager.GetTree(), fileIdString, constants.AllowedDiscreteFileExtensions...)
		if discreteFileInTree.Id == "" {
			ErrorHandler(w, fmt.Sprintf("Could not get resource %v", fileIdString), http.StatusBadRequest)
			return
		}

		filePathOnHardDisk := path.Join(config.RootPath, discreteFileInTree.Path)
		fileBytes, err := os.ReadFile(filePathOnHardDisk)
		if err != nil {
			fmt.Println(err.Error())
			ErrorHandler(w, fmt.Sprintf("Could not get resource %v", fileIdString), http.StatusBadRequest)
			return
		}

		w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\"", discreteFileInTree.Name))
		mimeType := utilities.GetContentTypeHeaderMimeType(discreteFileInTree)
		w.Header().Add("Content-Type", mimeType)
		w.Write(fileBytes)
	}

}
