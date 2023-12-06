package usecases

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
	"videos-with-subtitle-player/models"
	"videos-with-subtitle-player/router"
)

// TODO getSubtitleFileUseCase
// TODO => OUTSOURCE SAME LOGIC?

// initialize path as const on root level of the file to prevent circular dependencies
const GetAudioFileUseCasePath = `\/(.)+(\.mp3(\.vtt){0,1})$`

var GetAudioFileUseCaseRoute = router.Route{
	Path:    GetAudioFileUseCasePath,
	Method:  http.MethodGet,
	Handler: getAudioFileHandler,
}

func getAudioFileHandler(w http.ResponseWriter, r *http.Request, quit chan<- bool) {
	filePath := getFilePathFromUrl(r.URL.Path)
	audioFile, err := findAudioFileInFileTree(filePath)
	if err != nil {
		fmt.Println(err.Error())
		router.ErrorHandler(w, fmt.Sprintf("Could not get resource %v", filePath), http.StatusBadRequest)
		return
	}

	fileBytes, err := os.ReadFile(fmt.Sprintf(audioFile.Path))
	if err != nil {
		fmt.Println(err.Error())
		router.ErrorHandler(w, fmt.Sprintf("Could not get resource %v", filePath), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\"", audioFile.Name))
	w.Write(fileBytes)
	quit <- true
}

func getFilePathFromUrl(urlPath string) string {
	filePathRegexp := regexp.MustCompile(GetAudioFileUseCasePath)
	filePath := filePathRegexp.FindString(urlPath)
	filePathWithoutFilePathPrefix := strings.Replace(filePath, "/file", "", 1)
	return filePathWithoutFilePathPrefix
}

func findAudioFileInFileTree(filePath string) (models.FileTreeItem, error) {
	rootPath := os.Getenv("ROOT_PATH")
	// TODO NEW FUNCTION RECURSIVELY
	fileTree := GetFileTree(rootPath)
	fmt.Println(fileTree)
	return models.FileTreeItem{}, nil
}
