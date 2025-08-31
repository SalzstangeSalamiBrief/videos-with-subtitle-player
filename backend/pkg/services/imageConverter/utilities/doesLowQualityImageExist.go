package utilities

import (
	"os"
	"path/filepath"
)

func DoesLowQualityImageExist(rootPath string, sourceImagePath string) bool {
	if filepath.Ext(sourceImagePath) == "" {
		return false
	}

	lowQualityImagePath := GetLowQualityImagePath(sourceImagePath)
	absolutePath := filepath.Join(rootPath, lowQualityImagePath)
	_, err := os.Stat(absolutePath)
	return !os.IsNotExist(err)
}
