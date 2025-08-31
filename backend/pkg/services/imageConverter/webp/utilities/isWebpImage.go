package utilities

import (
	"backend/pkg/services/imageConverter/constants"
	"path/filepath"
)

func IsWebpImage(relativeImagePath string) bool {
	return filepath.Ext(relativeImagePath) == constants.WebpExtension
}
