package usecases

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"videos-with-subtitle-player/models"
	"videos-with-subtitle-player/router"
	directoryTree "videos-with-subtitle-player/services/directoryTree"
)

const GetAudioFileUseCasePath = `\/file\/([0-9A-Fa-f]{8}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{12})$$`

var GetAudioFileUseCaseRoute = router.Route{
	Path:    GetAudioFileUseCasePath,
	Method:  http.MethodGet,
	Handler: getAudioFileHandler,
}

func getAudioFileHandler(w http.ResponseWriter, r *http.Request, quit chan<- bool) {
	rootPath := os.Getenv("ROOT_PATH")
	fileIdString := strings.TrimPrefix(r.URL.Path, "/api/file/")
	audioFileInTree := getFileById(fileIdString)
	if audioFileInTree.Id == "" {
		router.ErrorHandler(w, fmt.Sprintf("Could not get resource %v", fileIdString), http.StatusBadRequest)
		return
	}
	filePathOnHardDist := path.Join(rootPath, audioFileInTree.Path)
	fileBytes, err := os.ReadFile(filePathOnHardDist)
	if err != nil {
		fmt.Println(err.Error())
		router.ErrorHandler(w, fmt.Sprintf("Could not get resource %v", fileIdString), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\"", audioFileInTree.Name))
	w.Write(fileBytes)
	quit <- true
}

func getFileById(fileIdString string) models.FileTreeItem {
	var file models.FileTreeItem
	for _, fileTreeItem := range directoryTree.FileTreeItems {
		isMatch := fileTreeItem.Id == fileIdString
		if isMatch {
			file = fileTreeItem
			break
		}
	}

	return file
}
