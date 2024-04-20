package lib

import (
	"path"
	"strings"
)

func GetFileNameWithoutExtension(filename string) string {
	fileExtension := path.Ext(filename)
	fileNameWithoutExtension := strings.Replace(filename, fileExtension, "", 1)
	// used if two file names are chained e. g. .mp3.vtt
	fileExtension = path.Ext(fileNameWithoutExtension)
	if fileExtension != "" {
		fileNameWithoutExtension = strings.Replace(fileNameWithoutExtension, fileExtension, "", 1)
	}

	return fileNameWithoutExtension
}
