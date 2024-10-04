package logic

import (
	"path"
	"strings"
)

func GetFilenameWithoutExtension(filename string) string {
	fileExtension := path.Ext(filename)
	fileNameWithoutExtension := strings.Replace(filename, fileExtension, "", 1)

	isSubtitleFile := fileExtension == ".vtt"
	if isSubtitleFile {
		fileExtension = path.Ext(fileNameWithoutExtension)
		if fileExtension != "" {
			fileNameWithoutExtension = strings.Replace(fileNameWithoutExtension, fileExtension, "", 1)
		}
	}

	return fileNameWithoutExtension
}
