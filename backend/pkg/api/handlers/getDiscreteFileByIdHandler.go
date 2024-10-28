package handlers

import (
	"backend/pkg/services/fileTreeManager/constants"
	"backend/pkg/utilities"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
)

func GetDiscreteFileHandler(w http.ResponseWriter, r *http.Request) {
	rootPath := os.Getenv("ROOT_PATH")
	fileIdString := strings.TrimPrefix(r.URL.Path, "/api/file/discrete/")
	discreteFileInTree := utilities.GetFileByIdAndExtension(fileIdString, constants.AllowedDiscreteFileExtensions...)
	if discreteFileInTree.Id == "" {
		ErrorHandler(w, fmt.Sprintf("Could not get resource %v", fileIdString), http.StatusBadRequest)
		return
	}

	filePathOnHardDisk := path.Join(rootPath, discreteFileInTree.Path)
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
