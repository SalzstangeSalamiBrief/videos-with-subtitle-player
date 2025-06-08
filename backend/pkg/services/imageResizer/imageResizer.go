package imageResizer

import (
	"backend/pkg/utilities"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

const resizeImageWidth = "640x"
const resizeFileSuffix = "_resize"

func Resize(sourceImagePath string) (resizeImagePath string, err error) {
	resizeImagePath = getResizeImagePath(sourceImagePath)
	doesResizeFilePathExist := utilities.DoesFileExist(resizeImagePath)
	if doesResizeFilePathExist {
		log.Default().Printf("File %s already has a resized version", sourceImagePath)
		return resizeImagePath, nil
	}

	err = executeResize(sourceImagePath, resizeImagePath)
	return resizeImagePath, err
}

func IsResizeFileName(sourceImagePath string) bool {
	return strings.Contains(filepath.Base(sourceImagePath), resizeFileSuffix)
}

func getResizeImagePath(sourceImagePath string) string {
	inputFileName, inputFileExtension := getFilenameAndExtensionParts(sourceImagePath)
	resizeImageFileName := getResizeImageName(inputFileName, inputFileExtension)
	return addPathToResizeImage(sourceImagePath, resizeImageFileName)
}

func getFilenameAndExtensionParts(sourcePath string) (name string, extension string) {
	name = ""
	extension = ""

	if len(sourcePath) == 0 {
		return name, extension
	}

	extension = filepath.Ext(sourcePath)
	name = strings.TrimSuffix(filepath.Base(sourcePath), extension)
	return name, extension
}

func getResizeImageName(name string, extension string) string {
	if name == "" || extension == "" {
		return ""
	}

	return fmt.Sprintf("%s%s%s", name, resizeFileSuffix, extension)
}

func addPathToResizeImage(inputFilePath string, resizeImageFileName string) string {
	if inputFilePath == "" || resizeImageFileName == "" {
		return ""
	}

	return filepath.Join(filepath.Dir(inputFilePath), resizeImageFileName)
}

func executeResize(inputFilePath string, resizeFilePath string) error {
	if _, err := exec.LookPath("magick"); err != nil {
		return fmt.Errorf("ImageMagick 'magick' command not found in PATH: %w", err)
	}

	command := exec.Command("magick", filepath.Clean(inputFilePath), "-resize", resizeImageWidth, filepath.Clean(resizeFilePath))
	return command.Run()
}
