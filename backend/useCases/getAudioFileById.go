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
	audioFileInTree := utilities.GetFileByIdAndExtension(fileIdString, ".wav", ".mp3", ".mp4")
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

	fileSize, err := getFileSize(file)
	if err != nil {
		fmt.Println(err.Error())
		router.ErrorHandler(w, fmt.Sprintf("Could not get resource '%v'", fileIdString), http.StatusInternalServerError)
		return
	}

	rangeHeaderWithPrefix := r.Header.Get("Range")
	start, end := utilities.GetRequestedRangesFromHeaderField(utilities.GetRequestRangesInput{rangeHeaderWithPrefix, chunkSize, fileSize})
	if start == 0 && end == 0 {
		router.ErrorHandler(w, fmt.Sprintf("The request does not contain a range header for file '%v'", fileIdString), http.StatusBadRequest)
		return
	}

	_, err = file.Seek(start, io.SeekStart)
	if err != nil {
		fmt.Println(err.Error())
		router.ErrorHandler(w, fmt.Sprintf("Could not get resource %v", fileIdString), http.StatusInternalServerError)
		return
	}

	addPartialContentHeader(w, start, end, fileSize)
	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\"", audioFileInTree.Name))

	mimeType := utilities.GetContentTypeHeaderMimeType(audioFileInTree)
	w.Header().Add("Content-Type", mimeType)
	w.WriteHeader(http.StatusPartialContent)
	io.CopyN(w, file, end-start)
}

func getFileSize(file *os.File) (int64, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return -1, err
	}

	fileSize := fileInfo.Size()
	return fileSize, nil
}

func addPartialContentHeader(w http.ResponseWriter, start int64, end int64, size int64) {
	w.Header().Add("Content-Range", fmt.Sprintf("bytes %v-%v/%v", start, end, size))
	w.Header().Add("Accept-Ranges", "bytes")
	w.Header().Add("Content-length", fmt.Sprint(end-start))
}
