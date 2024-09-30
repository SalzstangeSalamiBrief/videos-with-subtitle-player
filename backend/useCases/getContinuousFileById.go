package usecases

import (
	lib "backend/lib/utilities/models"
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
const GetContinuousFileUseCasePath = `\/file\/continuous\/([0-9A-Fa-f]{8}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{4}[-]?[0-9A-Fa-f]{12})$$`

var GetContinuousFileUseCaseRoute = router.Route{
	Path:    GetContinuousFileUseCasePath,
	Method:  http.MethodGet,
	Handler: getContinuousFileHandler,
}

func getContinuousFileHandler(w http.ResponseWriter, r *http.Request) {
	rootPath := os.Getenv("ROOT_PATH")
	fileIdString := strings.TrimPrefix(r.URL.Path, "/api/file/continuous/")
	continuousFileInTree := usecases.GetFileByIdAndExtension(fileIdString, lib.AllowedContinuousFileExtensions...)
	if continuousFileInTree.Id == "" {
		router.ErrorHandler(w, fmt.Sprintf("Could not get resource %v", fileIdString), http.StatusBadRequest)
		return
	}

	filePathOnHardDisk := path.Join(rootPath, continuousFileInTree.Path)
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
	start, end := usecases.GetRequestedRangesFromHeaderField(usecases.GetRequestRangesInput{RangeHeaderWithPrefix: rangeHeaderWithPrefix, ChunkSize: chunkSize, FileSize: fileSize})
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
	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\"", continuousFileInTree.Name))

	mimeType := usecases.GetContentTypeHeaderMimeType(continuousFileInTree)
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
