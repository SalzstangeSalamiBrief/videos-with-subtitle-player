package utilities

import "os"

func DoesFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return false
	}

	return !os.IsNotExist(err)
}
