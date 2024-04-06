package lib

import (
	"backend/models/enums"
	"fmt"
	"path"
	"regexp"
	"strings"
)

// TODO RENAME FILE
func getFileType(fileName string) (enums.FileType, error) {
	extension := path.Ext(fileName)
	switch extension {
	case ".mp3":
		return enums.AUDIO, nil
	case ".mp4":
		return enums.VIDEO, nil
	case ".wav":
		return enums.AUDIO, nil
	case ".vtt":
		return enums.SUBTITLE, nil
	default:
		return enums.UNKNOWN, fmt.Errorf("unknown file type for file '%v'", fileName)
	}
}

func getFolderPath(path string) string {
	pathWithoutRoot := strings.Replace(path, rootPath, "", 1)
	regexpToAddMatchingSeparators := regexp.MustCompile(`\\+`)
	pathWithSeparators := regexpToAddMatchingSeparators.ReplaceAllString(pathWithoutRoot, "/")
	return pathWithSeparators
}

func getFileName(path string) string {
	regexpToAddMatchingSeparators := regexp.MustCompile(`\\+`)
	pathParts := regexpToAddMatchingSeparators.Split(path, -1)
	fileName := pathParts[len(pathParts)-1]
	return fileName
}

func getFileNameWithoutExtension(filename string) string {
	fileExtension := path.Ext(filename)
	fileNameWithoutExtension := strings.Replace(filename, fileExtension, "", 1)
	// used if two file names are chained e. g. .mp3.vtt
	fileExtension = path.Ext(fileNameWithoutExtension)
	if fileExtension != "" {
		fileNameWithoutExtension = strings.Replace(fileNameWithoutExtension, fileExtension, "", 1)
	}

	return fileNameWithoutExtension
}
