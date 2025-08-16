package imageHandlerSources

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

type MagickImageHandler struct {
	lowQualityFileSuffix     string
	lowQualityConversionArgs []string
}

func NewMagickImageHandler(lowQualityFileSuffix string) *MagickImageHandler {
	imageMagickCommands := []ImageMagickCommand{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "10"}}
	return &MagickImageHandler{
		lowQualityFileSuffix:     lowQualityFileSuffix,
		lowQualityConversionArgs: convertImageMagickCommandsArrayToArgumentsArray(imageMagickCommands),
	}
}

func (imageHandler MagickImageHandler) ReduceImageQuality(sourceImagePath string) (lowQualityImagePath string, err error) {
	if imageHandler.lowQualityConversionArgs == nil || len(imageHandler.lowQualityConversionArgs) == 0 {
		log.Panic("imageQualityReducer is not properly initialized. Please use the InitializeMagickArgs function before using it")
	}

	lowQualityImagePath = getLowQualityImagePath(sourceImagePath, imageHandler.lowQualityFileSuffix)
	doesLowQualityFilePathExist := utilities.DoesFileExist(lowQualityImagePath)
	if doesLowQualityFilePathExist {
		log.Default().Printf("File %s already has a low quality version", sourceImagePath)
		return lowQualityImagePath, nil
	}

	err = executeReduceImageQuality(sourceImagePath, lowQualityImagePath, imageHandler.lowQualityConversionArgs)
	return lowQualityImagePath, err
}

func (imageHandler MagickImageHandler) IsLowQualityFile(sourcePath string) bool {
	return strings.Contains(filepath.Base(sourcePath), imageHandler.lowQualityFileSuffix)
}

func getLowQualityImagePath(sourceImagePath string, lowQualityFileSuffix string) string {
	inputFileName, inputFileExtension := getFilenameAndExtensionParts(sourceImagePath)
	lowQualityImageFileName := getLowQualityImageName(inputFileName, inputFileExtension, lowQualityFileSuffix)
	return addPathToLowQualityImage(sourceImagePath, lowQualityImageFileName)
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

func getLowQualityImageName(name string, extension string, lowQualityFileSuffix string) string {
	if name == "" || extension == "" {
		return ""
	}

	return fmt.Sprintf("%s%s%s", name, lowQualityFileSuffix, extension)
}

func addPathToLowQualityImage(sourceFilePath string, lowQualityImageFileName string) string {
	if sourceFilePath == "" || lowQualityImageFileName == "" {
		return ""
	}

	return filepath.Join(filepath.Dir(sourceFilePath), lowQualityImageFileName)
}

func executeReduceImageQuality(sourceFilePath string, lowQualityFilePath string, arguments []string) error {
	if _, err := exec.LookPath("magick"); err != nil {
		return fmt.Errorf("ImageMagick 'magick' command not found in PATH: %w", err)
	}

	log.Default().Printf("Start quality reducing process for source '%v'\n", sourceFilePath)
	command := exec.Command("magick", filepath.Clean(sourceFilePath))
	command.Args = append(command.Args, arguments...)
	command.Args = append(command.Args, filepath.Clean(lowQualityFilePath))
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
