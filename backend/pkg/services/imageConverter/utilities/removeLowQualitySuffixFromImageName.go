package utilities

import (
	"backend/pkg/services/imageConverter/constants"
	"strings"
)

func RemoveLowQualitySuffixFromImageName(relativePath string) string {
	return strings.Replace(relativePath, constants.LowQualityFileSuffix, "", 1)
}
