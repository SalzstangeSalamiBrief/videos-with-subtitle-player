package directorytree

import (
	"path"
	"regexp"
	"strings"
)

func getFolderPath(path string) string {
	pathWithoutRoot := strings.Replace(path, rootPath, "", 1)
	regexpToAddMatchingSeparators := regexp.MustCompile(`\\+`)
	pathWithSeparators := regexpToAddMatchingSeparators.ReplaceAllString(pathWithoutRoot, "/")
	return pathWithSeparators
}

func getFileName(path string) string {
	regexpToAddmatchingSeparators := regexp.MustCompile(`\\+`)
	pathParts := regexpToAddmatchingSeparators.Split(path, -1)
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
