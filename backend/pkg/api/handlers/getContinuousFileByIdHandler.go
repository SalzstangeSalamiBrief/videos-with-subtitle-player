package handlers

import (
	"backend/pkg/services/fileTreeManager/constants"
	"backend/pkg/utilities"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

const chunkSize = 1 * 1024 * 1024 // 3mb

func GetContinuousFileByIdHandler(w http.ResponseWriter, r *http.Request) {
	rootPath := os.Getenv("ROOT_PATH")
	fileIdString := strings.TrimPrefix(r.URL.Path, "/api/file/continuous/")
	continuousFileInTree := utilities.GetFileByIdAndExtension(fileIdString, constants.AllowedContinuousFileExtensions...)
	if continuousFileInTree.Id == "" {
		ErrorHandler(w, fmt.Sprintf("Could not get resource %v", fileIdString), http.StatusBadRequest)
		return
	}

	filePathOnHardDisk := path.Join(rootPath, continuousFileInTree.Path)
	file, err := os.Open(filePathOnHardDisk)
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
		ErrorHandler(w, fmt.Sprintf("Could not get resource '%v'", fileIdString), http.StatusInternalServerError)
		return
	}

	fileSize, err := getFileSize(file)
	if err != nil {
		fmt.Println(err.Error())
		ErrorHandler(w, fmt.Sprintf("Could not get resource '%v'", fileIdString), http.StatusInternalServerError)
		return
	}

	rangeHeaderWithPrefix := r.Header.Get("Range")
	start, end := utilities.GetRequestedRangesFromHeaderField(utilities.GetRequestRangesInput{RangeHeaderWithPrefix: rangeHeaderWithPrefix, ChunkSize: chunkSize, FileSize: fileSize})
	if start == 0 && end == 0 {
		ErrorHandler(w, fmt.Sprintf("The request does not contain a range header for file '%v'", fileIdString), http.StatusBadRequest)
		return
	}

	_, err = file.Seek(start, io.SeekStart)
	if err != nil {
		fmt.Println(err.Error())
		ErrorHandler(w, fmt.Sprintf("Could not get resource %v", fileIdString), http.StatusInternalServerError)
		return
	}

	addPartialContentHeader(w, start, end, fileSize)
	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\"", continuousFileInTree.Name))

	mimeType := utilities.GetContentTypeHeaderMimeType(continuousFileInTree)
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
