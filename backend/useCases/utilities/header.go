package utilities

import (
	"backend/lib"
	"backend/models"
	"fmt"
	"mime"
	"net/http"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

func AddContentTypeHeader(w http.ResponseWriter, selectedFile models.FileTreeItem) {
	ext := filepath.Ext(selectedFile.Path)
	mimeType := mime.TypeByExtension(ext)
	if ext == ".vtt" {
		mimeType = "text/vtt"
	}
	w.Header().Add("Content-Type", mimeType)
}

func AddPartialContentHeader(w http.ResponseWriter, start int64, end int64, size int64) {
	w.Header().Add("Content-Range", fmt.Sprintf("bytes %v-%v/%v", start, end, size))
	w.Header().Add("Accept-Ranges", "bytes")
	w.Header().Add("Content-length", fmt.Sprint(end-start))
}

func GetFileByIdAndExtension(id string, allowedExtension ...string) models.FileTreeItem {
	var file models.FileTreeItem
	for _, fileTreeItem := range lib.FileTreeItems {
		isMatch := fileTreeItem.Id == id
		if !isMatch {
			continue
		}

		ext := filepath.Ext(fileTreeItem.Path)
		doesExtensionMatch := slices.Contains(allowedExtension, ext)
		if doesExtensionMatch {
			file = fileTreeItem
			break
		}
	}

	return file
}

func GetRequestedRangesFromHeader(r *http.Request, chunkSize int64) (start int64, end int64) {
	rangeHeaderWithPrefix := r.Header.Get("Range")
	if rangeHeaderWithPrefix == "" {
		return 0, chunkSize
	}

	rangeHeaderWithoutPrefix := strings.TrimPrefix(rangeHeaderWithPrefix, "bytes=")
	rangeParts := strings.Split(rangeHeaderWithoutPrefix, "-")
	if len(rangeParts) != 2 {
		return 0, 0
	}

	start, err := strconv.ParseInt(rangeParts[0], 10, 64)
	if err != nil {
		fmt.Println("start", err.Error())
		return 0, 0
	}

	end = 0

	if rangeParts[1] == "" {
		return start, start + chunkSize
	}

	end, err = strconv.ParseInt(rangeParts[1], 10, 64)
	if err != nil {
		fmt.Println("end", err.Error())
		return 0, 0
	}

	if end == 0 {
		end = start + chunkSize
	}

	return start, end
}
