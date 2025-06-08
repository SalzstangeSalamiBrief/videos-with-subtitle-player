package utilities

import "os"

func DoesFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}

	return !os.IsNotExist(err)
}
