package utilities

import (
	"backend/pkg/services/imageConverter/constants"
	"fmt"
	"path/filepath"
	"strings"
)

func GetLowQualityImagePath(inputRelativeFilePath string) string {
	extension := filepath.Ext(inputRelativeFilePath)
	newFileSuffix := fmt.Sprintf("%s%s", constants.LowQualityFileSuffix, constants.WebpExtension)
	outputRelativeFilePath := strings.Replace(inputRelativeFilePath, extension, newFileSuffix, -1)
	return outputRelativeFilePath
}
