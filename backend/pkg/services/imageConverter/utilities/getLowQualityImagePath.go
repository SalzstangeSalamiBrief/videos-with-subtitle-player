package utilities

import (
	"backend/pkg/services/imageConverter/constants"
	"fmt"
	"path/filepath"
	"strings"
)

func GetLowQualityImagePath(inputRelativeFilePath string) string {
	extension := filepath.Ext(inputRelativeFilePath)
	if extension == "" {
		return ""
	}

	newFileSuffix := fmt.Sprintf("%s%s", constants.LowQualityFileSuffix, extension)
	outputRelativeFilePath := strings.Replace(inputRelativeFilePath, extension, newFileSuffix, -1)
	return outputRelativeFilePath
}
