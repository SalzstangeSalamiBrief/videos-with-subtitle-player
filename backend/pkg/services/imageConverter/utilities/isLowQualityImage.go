package utilities

import (
	"backend/pkg/services/imageConverter/constants"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
)

var lowQualityImageSuffix = regexp.MustCompile(fmt.Sprintf("%s$", constants.LowQualityFileSuffix))

func IsLowQualityImage(relativeImagePath string) bool {
	fileName := filepath.Base(relativeImagePath)
	fileNameWithoutExtension := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	isLowQualityImage := lowQualityImageSuffix.MatchString(fileNameWithoutExtension)
	return isLowQualityImage
}
