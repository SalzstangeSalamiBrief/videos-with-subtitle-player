package imageResizer

import (
	"fmt"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

const resizeImageWidth = "640x"
const resizeFileSuffix = "_resize"

func Resize(inputFilePath string) error {
	inputFileName, inputFileExtension := getFilenameAndExtensionParts(inputFilePath)
	resizeImageFileName := getResizeImageName(inputFileName, inputFileExtension)
	resizeImageFilePath := addPathToResizeImage(inputFilePath, resizeImageFileName)
	return executeResize(inputFileName, resizeImageFilePath)
}

func getFilenameAndExtensionParts(filePath string) (name string, extension string) {
	name = ""
	extension = ""

	if len(filePath) == 0 {
		return name, extension
	}

	var base = path.Base(filePath)
	extension = path.Ext(base)
	name = strings.TrimSuffix(base, extension)
	return name, extension
}

func getResizeImageName(name string, extension string) string {
	if name == "" || extension == "" {
		return ""
	}

	return fmt.Sprintf("%s%s%s", name, resizeFileSuffix, extension)
}

func addPathToResizeImage(inputFilePath string, resizeImageFileName string) string {
	return fmt.Sprintf("%v/%v", filepath.Dir(inputFilePath), resizeImageFileName)
}

func executeResize(inputFilePath string, resizeFilePath string) error {
	command := exec.Command("magick", inputFilePath, "-resize", resizeImageWidth, resizeFilePath)
	return command.Run()
}
