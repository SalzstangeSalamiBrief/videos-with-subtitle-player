package utilities

import (
	"backend/pkg/services/imageConverter/constants"
	"fmt"
	"path/filepath"
	"strings"
)

func GetLowQualityImageName(sourceImagePath string) string {
	if filepath.Ext(sourceImagePath) == "" {
		return ""
	}

	if sourceImagePath == "" {
		return ""
	}

	return fmt.Sprintf("%s%s%s", strings.TrimSuffix(filepath.Base(sourceImagePath), filepath.Ext(sourceImagePath)), constants.LowQualityFileSuffix, filepath.Ext(sourceImagePath))
}
