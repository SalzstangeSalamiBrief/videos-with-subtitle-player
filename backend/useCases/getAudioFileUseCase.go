package usecases

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"videos-with-subtitle-player/models"
	"videos-with-subtitle-player/router"
)

// initialize path as const on root level of the file to prevent circular dependencies
const GetAudioFileUseCasePath = `\/file\/(.)+(\.mp3(\.vtt){0,1})$`

// this use case can be used to get .mp3 oder .vtt files
var GetAudioFileUseCaseRoute = router.Route{
	Path:    GetAudioFileUseCasePath,
	Method:  http.MethodGet,
	Handler: getAudioFileHandler,
}

func getAudioFileHandler(w http.ResponseWriter, r *http.Request, quit chan<- bool) {
	rootPath := os.Getenv("ROOT_PATH")
	filePath := getFilePathFromUrl(r.URL.Path)
	audioFileInTree, err := findAudioFileInFileTree(filePath, rootPath)
	if err != nil {
		fmt.Println(err.Error())
		router.ErrorHandler(w, fmt.Sprintf("Could not get resource %v", filePath), http.StatusBadRequest)
		return
	}

	filePathOnHardDist := path.Join(rootPath, audioFileInTree.Path)
	fileBytes, err := os.ReadFile(filePathOnHardDist)
	if err != nil {
		fmt.Println(err.Error())
		router.ErrorHandler(w, fmt.Sprintf("Could not get resource %v", filePath), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\"", audioFileInTree.Name))
	w.Write(fileBytes)
	quit <- true
}

func getFilePathFromUrl(urlPath string) string {
	filePathRegexp := regexp.MustCompile(GetAudioFileUseCasePath)
	filePath := filePathRegexp.FindString(urlPath)
	filePathWithoutFilePathPrefix := strings.Replace(filePath, "/file", "", 1)
	return filePathWithoutFilePathPrefix
}

func findAudioFileInFileTree(filePath string, rootPath string) (models.FileTreeItem, error) {

	flatFileTree := GetFlatFileTree(rootPath)

	var err error
	var result models.FileTreeItem
	for _, fileTreeItem := range flatFileTree {
		isMatch := strings.Contains(fileTreeItem.Path, filePath)
		if isMatch == true {
			result = fileTreeItem
			break
		}
	}

	if result.Path == "" {
		err = fmt.Errorf("Could not find file '%v' in file tree", filePath)
	}

	return result, err
}
