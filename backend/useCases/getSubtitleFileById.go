package usecases

import (
	"backend/router"
	"backend/useCases/utilities"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
)

const GetSubtitleFileUseCasePath = `\/file\/subtitle\/([0-9A-Fa-f]{8}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{12})$$`

var GetSubtitleFileUseCaseRoute = router.Route{
	Path:    GetSubtitleFileUseCasePath,
	Method:  http.MethodGet,
	Handler: getSubtitleFileHandler,
}

func getSubtitleFileHandler(w http.ResponseWriter, r *http.Request) {
	rootPath := os.Getenv("ROOT_PATH")
	fileIdString := strings.TrimPrefix(r.URL.Path, "/api/file/subtitle/")
	subtitleFileInTree := utilities.GetFileByIdAndExtension(fileIdString, ".vtt")
	if subtitleFileInTree.Id == "" {
		router.ErrorHandler(w, fmt.Sprintf("Could not get resource %v", fileIdString), http.StatusBadRequest)
		return
	}

	filePathOnHardDisk := path.Join(rootPath, subtitleFileInTree.Path)
	fileBytes, err := os.ReadFile(filePathOnHardDisk)
	if err != nil {
		fmt.Println(err.Error())
		router.ErrorHandler(w, fmt.Sprintf("Could not get resource %v", fileIdString), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\"", subtitleFileInTree.Name))
	mimeType := utilities.GetContentTypeHeaderMimeType(subtitleFileInTree)
	w.Header().Add("Content-Type", mimeType)
	w.Write(fileBytes)
}
