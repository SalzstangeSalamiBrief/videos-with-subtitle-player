package usecases

import (
	"backend/router"
	"backend/useCases/utilities"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

const chunkSize = 1 * 1024 * 1024 // 3mb
const GetAudioFileUseCasePath = `\/file\/audio\/([0-9A-Fa-f]{8}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{12})$$`

var GetAudioFileUseCaseRoute = router.Route{
	Path:    GetAudioFileUseCasePath,
	Method:  http.MethodGet,
	Handler: getAudioFileHandler,
}

func getAudioFileHandler(w http.ResponseWriter, r *http.Request) {
	rootPath := os.Getenv("ROOT_PATH")
	fileIdString := strings.TrimPrefix(r.URL.Path, "/api/file/audio/")
	audioFileInTree := utilities.GetFileByIdAndExtension(fileIdString, ".wav", ".mp3")
	if audioFileInTree.Id == "" {
		router.ErrorHandler(w, fmt.Sprintf("Could not get resource %v", fileIdString), http.StatusBadRequest)
		return
	}

	filePathOnHardDisk := path.Join(rootPath, audioFileInTree.Path)
	file, err := os.Open(filePathOnHardDisk)
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
		router.ErrorHandler(w, fmt.Sprintf("Could not get resource '%v'", fileIdString), http.StatusInternalServerError)
		return
	}

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
		router.ErrorHandler(w, fmt.Sprintf("Could not get resource '%v'", fileIdString), http.StatusInternalServerError)
		return
	}

	fileSize := fileInfo.Size()
	start, end := utilities.GetRequestedRangesFromHeader(r, chunkSize)
	if start == 0 && end == 0 {
		router.ErrorHandler(w, fmt.Sprintf("The request does not contain a range header for file '%v'", fileIdString), http.StatusBadRequest)
		return
	}

	if end > fileSize {
		end = fileSize
	}

	_, err = file.Seek(start, io.SeekStart)
	if err != nil {
		fmt.Println(err.Error())
		router.ErrorHandler(w, fmt.Sprintf("Could not get resource %v", fileIdString), http.StatusInternalServerError)
		return
	}

	utilities.AddPartialContentHeader(w, start, end, fileSize)
	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\"", audioFileInTree.Name))
	utilities.AddContentTypeHeader(w, audioFileInTree)
	w.WriteHeader(http.StatusPartialContent)
	io.CopyN(w, file, end-start)
}
