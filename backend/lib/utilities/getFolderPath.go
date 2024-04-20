package lib

import (
	"regexp"
	"strings"
)

func GetFolderPath(path string, rootPath string) string {
	pathWithoutRoot := strings.Replace(path, rootPath, "", 1)
	regexpToAddMatchingSeparators := regexp.MustCompile(`\\+`)
	pathWithSeparators := regexpToAddMatchingSeparators.ReplaceAllString(pathWithoutRoot, "/")
	return pathWithSeparators
}
