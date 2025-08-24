package utilities

import (
	"os"
	"path/filepath"
)

func DoesLowQualityImageExist(rootPath string, sourceImagePath string) bool {
	lowQualityImagePath := GetLowQualityImagePath(sourceImagePath)
	absolutePath := filepath.Join(rootPath, lowQualityImagePath)
	_, err := os.Stat(absolutePath)
	return !os.IsNotExist(err)
}
