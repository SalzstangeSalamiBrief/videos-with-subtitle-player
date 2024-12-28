package handlers

import (
	"backend/pkg/errors"
	"backend/pkg/services/config"
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
	fileIdString := strings.TrimPrefix(r.URL.Path, "/api/file/continuous/")
	continuousFileInTree := utilities.GetFileByIdAndExtension(fileIdString, constants.AllowedContinuousFileExtensions...)
	if continuousFileInTree.Id == "" {
		ErrorHandler(w, &errors.FileNotFoundError{Id: fileIdString})
		return
	}

	filePathOnHardDisk := path.Join(config.AppConfiguration.RootPath, continuousFileInTree.Path)
	file, err := os.Open(filePathOnHardDisk)
	defer file.Close()
	if err != nil {
		fmt.Println(err.Error())
		ErrorHandler(w, &errors.FileNotFoundError{Id: fileIdString})
		return
	}

	fileSize, err := getFileSize(file)
	if err != nil {
		ErrorHandler(w, &errors.OsError{Id: fileIdString, InnerError: err.Error()})
		return
	}

	rangeHeaderWithPrefix := r.Header.Get("Range")
	start, end := utilities.GetRequestedRangesFromHeaderField(utilities.GetRequestRangesInput{RangeHeaderWithPrefix: rangeHeaderWithPrefix, ChunkSize: chunkSize, FileSize: fileSize})
	if start == 0 && end == 0 {
		ErrorHandler(w, &errors.MissingHeaderError{Id: fileIdString, Header: "Range"})
		return
	}

	_, err = file.Seek(start, io.SeekStart)
	if err != nil {
		fmt.Println(err.Error())
		ErrorHandler(w, &errors.OsError{Id: fileIdString, InnerError: err.Error()})
		return
	}

	addPartialContentHeader(w, start, end, fileSize)
	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\"", continuousFileInTree.Name))

	mimeType := continuousFileInTree.GetMimeType()
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
