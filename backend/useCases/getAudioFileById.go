package usecases

import (
	"backend/models"
	"backend/router"
	directorytree "backend/services/directoryTree"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
)

const GetAudioFileUseCasePath = `\/file\/([0-9A-Fa-f]{8}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{12})$$`

var GetAudioFileUseCaseRoute = router.Route{
	Path:    GetAudioFileUseCasePath,
	Method:  http.MethodGet,
	Handler: getAudioFileHandler,
}

func getAudioFileHandler(w http.ResponseWriter, r *http.Request) {
	rootPath := os.Getenv("ROOT_PATH")
	fileIdString := strings.TrimPrefix(r.URL.Path, "/api/file/")
	audioFileInTree := getFileById(fileIdString)
	if audioFileInTree.Id == "" {
		router.ErrorHandler(w, fmt.Sprintf("Could not get resource %v", fileIdString), http.StatusBadRequest)
		return
	}

	filePathOnHardDisk := path.Join(rootPath, audioFileInTree.Path)
	fileBytes, err := os.ReadFile(filePathOnHardDisk)
	if err != nil {
		fmt.Println(err.Error())
		router.ErrorHandler(w, fmt.Sprintf("Could not get resource %v", fileIdString), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\"", audioFileInTree.Name))
	addContentTypeHeader(w, audioFileInTree)
	w.Write(fileBytes)
}

func getFileById(fileIdString string) models.FileTreeItem {
	var file models.FileTreeItem
	for _, fileTreeItem := range directorytree.FileTreeItems {
		isMatch := fileTreeItem.Id == fileIdString
		if isMatch {
			file = fileTreeItem
			break
		}
	}

	return file
}

func addContentTypeHeader(w http.ResponseWriter, selectedFile models.FileTreeItem) {
	contentTypeToAdd := "audio/mpeg3"
	isSubtitleFile := strings.HasSuffix(selectedFile.Path, ".vtt")
	if isSubtitleFile == true {
		contentTypeToAdd = "text/vtt"
	}

	w.Header().Add("Content-Type", contentTypeToAdd)
}
