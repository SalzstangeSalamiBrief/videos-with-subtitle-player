package handlers

import (
	"backend/internal/problemDetailsErrors"
	"backend/pkg/constants"
	"backend/pkg/services/fileTreeManager"
	"backend/pkg/utilities"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

const chunkSize = 1 * 1024 * 1024 // 1mb

type ContinuousFileByIdHandlerConfiguration struct {
	RootPath string
	*fileTreeManager.FileTreeManager
}

func NewGetContinuousFileByIdHandler(configuration ContinuousFileByIdHandlerConfiguration) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fileIdString := strings.TrimPrefix(r.URL.Path, "/api/file/continuous/")
		if fileIdString == "" {
			log.Default().Println(fmt.Sprintf("[ContinuousFileByIdHandler]: Error while opening the file with id '%v'\n", fileIdString))
			problemDetailsErrors.NewBadRequestProblemDetails(fmt.Sprintf("The parameter 'fileId' is empty but required. Please provide an id\n")).SendErrorResponse(w)
			return
		}

		continuousFileInTree := utilities.GetFileByIdAndExtension(configuration.FileTreeManager.GetTree(), fileIdString, constants.AllowedContinuousFileExtensions...)
		if continuousFileInTree.FileId == "" {
			log.Default().Println(fmt.Sprintf("[ContinuousFileByIdHandler]: Error while opening the file with id '%v'\n", fileIdString))
			problemDetailsErrors.NewNotFoundProblemDetails(fmt.Sprintf("Could not find file with id '%v'\n", fileIdString)).SendErrorResponse(w)
			return
		}

		filePathOnHardDisk := path.Join(configuration.RootPath, continuousFileInTree.Path)
		file, err := os.Open(filePathOnHardDisk)
		defer file.Close()

		if err != nil {
			log.Default().Println(fmt.Sprintf("[ContinuousFileByIdHandler]: Error while opening the file with id '%v'\n", err.Error()))
			problemDetailsErrors.NewBadRequestProblemDetails(fmt.Sprintf("Could not find the file with id '%v': %v\n", fileIdString, err)).SendErrorResponse(w)
			return
		}

		fileSize, err := getFileSize(file)
		if err != nil {
			log.Default().Println(fmt.Sprintf("[ContinuousFileByIdHandler]: Error while getting size of the file with id '%v'\n", err.Error()))
			problemDetailsErrors.NewBadRequestProblemDetails(fmt.Sprintf("Could not find the file with id '%v': %v\n", fileIdString, err)).SendErrorResponse(w)
			return
		}

		rangeHeaderWithPrefix := r.Header.Get("Range")
		start, end := utilities.GetRequestedRangesFromHeaderField(utilities.GetRequestRangesInput{RangeHeaderWithPrefix: rangeHeaderWithPrefix, ChunkSize: chunkSize, FileSize: fileSize})
		if start == 0 && end == 0 {
			log.Default().Println(fmt.Sprintf("[ContinuousFileByIdHandler]: Error while the getting range header for file with id: '%v'\n", err.Error()))
			problemDetailsErrors.NewBadRequestProblemDetails(fmt.Sprintf("The request for file with id '%v' does not contain range header.\n", fileIdString)).SendErrorResponse(w)
			return
		}

		_, err = file.Seek(start, io.SeekStart)
		if err != nil {
			log.Default().Println(err.Error())
			problemDetailsErrors.NewInternalServerErrorProblemDetails(fmt.Sprintf("Failed to seek the file with id '%v' with the provided range header.\n", fileIdString)).SendErrorResponse(w)
			return
		}

		addPartialContentHeader(w, start, end, fileSize)
		w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\"", continuousFileInTree.Name))

		mimeType := utilities.GetContentTypeHeaderMimeType(continuousFileInTree)
		w.Header().Add("Content-Type", mimeType)
		w.WriteHeader(http.StatusPartialContent)
		io.CopyN(w, file, end-start)
	}
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
