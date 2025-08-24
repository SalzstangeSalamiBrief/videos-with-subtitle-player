package utilities

import (
	"backend/pkg/services/imageConverter/constants"
	"fmt"
	"path/filepath"
	"strings"
)

func IsLowQualityImage(relativeImagePath string) bool {
	return strings.Contains(filepath.Base(relativeImagePath), fmt.Sprintf("%s%s", constants.LowQualityFileSuffix, constants.WebpExtension))
}
