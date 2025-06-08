package utilities

import (
	"path/filepath"
	"regexp"
	"strings"
)

type GetFolderPathInput struct {
	Path     string
	RootPath string
}

func GetFolderPath(input GetFolderPathInput) string {
	// TODO REFACTOR use filepath.base, filepath.join...
	pathWithoutRoot := strings.Replace(input.Path, input.RootPath, "", 1)
	regexpToAddMatchingSeparators := regexp.MustCompile(`\\+`)
	pathWithSeparators := regexpToAddMatchingSeparators.ReplaceAllString(pathWithoutRoot, "/")
	return filepath.Join(pathWithSeparators)
}
