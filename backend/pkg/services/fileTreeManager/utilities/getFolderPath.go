package utilities

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type GetFolderPathInput struct {
	Path     string
	RootPath string
}

func GetFolderPath(input GetFolderPathInput) string {
	pathWithoutRoot := strings.Replace(input.Path, input.RootPath, "", 1)
	regexpToAddMatchingSeparators := regexp.MustCompile(`\\+`)
	pathWithSeparators := regexpToAddMatchingSeparators.ReplaceAllString(pathWithoutRoot, string(os.PathSeparator))
	return filepath.Join(pathWithSeparators)
}
