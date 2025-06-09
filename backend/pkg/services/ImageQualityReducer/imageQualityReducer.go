package ImageQualityReducer

import (
	"backend/pkg/utilities"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

type ImageMagickCommand struct {
	command string
	arg     string
}

const resizeFileSuffix = "_lowQuality"

var magickArgs []string

func ReduceImageQuality(sourceImagePath string) (resizeImagePath string, err error) {
	if magickArgs == nil || len(magickArgs) == 0 {
		imageMagickCommands := []ImageMagickCommand{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "10"}}
		magickArgs = convertImageMagickCommandsArrayToArgumentsArray(imageMagickCommands)
	}

	resizeImagePath = getResizeImagePath(sourceImagePath)
	doesResizeFilePathExist := utilities.DoesFileExist(resizeImagePath)
	if doesResizeFilePathExist {
		log.Default().Printf("File %s already has a resized version", sourceImagePath)
		return resizeImagePath, nil
	}

	err = executeReduceImageQuality(sourceImagePath, resizeImagePath, magickArgs)
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

func executeReduceImageQuality(inputFilePath string, resizeFilePath string, arguments []string) error {
	if _, err := exec.LookPath("magick"); err != nil {
		return fmt.Errorf("ImageMagick 'magick' command not found in PATH: %w", err)
	}

	command := exec.Command("magick", filepath.Clean(inputFilePath))
	command.Args = append(command.Args, arguments...)
	command.Args = append(command.Args, filepath.Clean(resizeFilePath))
	return command.Run()
}

func convertImageMagickCommandsArrayToArgumentsArray(commands []ImageMagickCommand) []string {
	result := []string{}
	for _, command := range commands {
		if command.command != "" {
			result = append(result, command.command)

		}

		if command.arg != "" {
			result = append(result, command.arg)

		}
	}

	return result
}
