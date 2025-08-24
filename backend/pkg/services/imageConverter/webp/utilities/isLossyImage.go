package utilities

import "path/filepath"

func IsLossyImage(relativeImagePath string) bool {
	extension := filepath.Ext(relativeImagePath)
	return extension == ".jpg" || extension == ".jpeg"
}
