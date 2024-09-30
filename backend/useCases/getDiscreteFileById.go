package usecases

import (
	lib "backend/lib/utilities/models"
	"backend/router"
	"backend/useCases/utilities"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
)

const GetDiscreteFileUseCasePath = `\/file\/discrete\/([0-9A-Fa-f]{8}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{12})$$`

var GetDiscreteFileUseCaseRoute = router.Route{
	Path:    GetDiscreteFileUseCasePath,
	Method:  http.MethodGet,
	Handler: getDiscreteFileHandler,
}

func getDiscreteFileHandler(w http.ResponseWriter, r *http.Request) {
	rootPath := os.Getenv("ROOT_PATH")
	fileIdString := strings.TrimPrefix(r.URL.Path, "/api/file/discrete/")
	discreteFileInTree := usecases.GetFileByIdAndExtension(fileIdString, lib.AllowedDiscreteFileExtensions...)
	if discreteFileInTree.Id == "" {
		router.ErrorHandler(w, fmt.Sprintf("Could not get resource %v", fileIdString), http.StatusBadRequest)
		return
	}

	filePathOnHardDisk := path.Join(rootPath, discreteFileInTree.Path)
	fileBytes, err := os.ReadFile(filePathOnHardDisk)
	if err != nil {
		fmt.Println(err.Error())
		router.ErrorHandler(w, fmt.Sprintf("Could not get resource %v", fileIdString), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\"", discreteFileInTree.Name))
	mimeType := usecases.GetContentTypeHeaderMimeType(discreteFileInTree)
	w.Header().Add("Content-Type", mimeType)
	w.Write(fileBytes)
}
