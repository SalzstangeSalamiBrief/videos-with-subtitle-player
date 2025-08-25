package main

import (
	"backend/internal/config"
	"backend/pkg/enums"
	"backend/pkg/services/imageConverter/constants"
	"backend/pkg/services/imageConverter/webp"
	"backend/pkg/utilities"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"slices"
	"strings"
)

type Command struct {
	command string
	arg     string
}

const imageSourceFolder = "./images/lossless"
const imageSource = "Thumbnail.jpg"

func main() {
	//webpPoc()
	initializedConfiguration := config.InitializeConfiguration()
	err := webp.ExecuteWebpConversion(webp.ExecuteWebpConversionConfiguration{RootPath: initializedConfiguration.RootPath, ShouldDeleteNonWebpImages: true})
	if err != nil {
		log.Fatal(err)
	}
}

func webpPoc() {
	commands := [][]Command{
		[]Command{{command: "-quality", arg: "10"}},
		[]Command{{command: "-quality", arg: "20"}},
		[]Command{{command: "-quality", arg: "50"}},
		[]Command{{command: "-quality", arg: "70"}},
		[]Command{{command: "-quality", arg: "80"}},
		[]Command{{command: "-quality", arg: "90"}},
		[]Command{{command: "-resize", arg: "640x"}},
		[]Command{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "10"}, {command: "-define", arg: "webp:method=0"}},
		[]Command{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "20"}, {command: "-define", arg: "webp:method=0"}},
		[]Command{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "50"}, {command: "-define", arg: "webp:method=0"}},
		[]Command{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "70"}, {command: "-define", arg: "webp:method=0"}},
		[]Command{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "80"}, {command: "-define", arg: "webp:method=0"}},
		[]Command{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "90"}, {command: "-define", arg: "webp:method=0"}},
		[]Command{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "10"}, {command: "-define", arg: "webp:method=6"}},
		[]Command{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "20"}, {command: "-define", arg: "webp:method=6"}},
		[]Command{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "50"}, {command: "-define", arg: "webp:method=6"}},
		[]Command{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "70"}, {command: "-define", arg: "webp:method=6"}},
		[]Command{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "80"}, {command: "-define", arg: "webp:method=6"}},
		[]Command{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "90"}, {command: "-define", arg: "webp:method=6"}},
	}

	for _, command := range commands {
		handleCommand(command)
	}
}

func handleCommand(commandArguments []Command) {
	fmt.Printf("Start for commands %v\n", commandArguments)
	extension := filepath.Ext(imageSource)
	filename := strings.TrimSuffix(imageSource, extension)
	inputImageRelativePath := path.Join(imageSourceFolder, imageSource)
	fmt.Println("inputImageRelativePath", inputImageRelativePath)
	resizedImageName := strings.Replace(fmt.Sprintf("%v%v%v", filename, getFilenameFromCommands(commandArguments), extension), extension, ".webp", -1)
	fmt.Println("resizedImageName", resizedImageName)
	resizedImagePath := path.Join(imageSourceFolder, resizedImageName)

	commandArgsArray := convertCommandsArrayToStringArray(commandArguments)
	args := []string{inputImageRelativePath}
	args = append(args, commandArgsArray...)
	args = append(args, resizedImagePath)
	command := exec.Command("magick", args...)
	fmt.Printf("Executing: magick %s\n", strings.Join(args, " "))
	err := command.Run()
	if err != nil {
		log.Printf("Command failed: %v", err)
		log.Printf("Full command: magick %s", strings.Join(args, " "))
		log.Fatal(err)
	}

	fmt.Printf("Finish for commands %v\n", commandArguments)
}

func getFilenameFromCommands(commands []Command) string {
	result := ""
	for _, command := range commands {
		result += strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf("%v%v", command.command, command.arg), "=", "_"), ":", "_")
	}

	return result
}

func convertCommandsArrayToStringArray(commands []Command) []string {
	var result []string
	for _, command := range commands {
		result = append(result, command.command)
		result = append(result, command.arg)
	}

	return result
}

const rootPath string = "F:\\visual_novel_stuff\\asmr\\server_resources"

func traversalPoc(rootPath string) []string {
	imageFilePaths := make([]string, 0)

	content, err := os.ReadDir(rootPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range content {
		if item.IsDir() {
			subContent := traversalPoc(filepath.Join(rootPath, item.Name()))
			if len(subContent) > 0 {
				imageFilePaths = append(imageFilePaths, subContent...)
			}
		}

		fileType := utilities.GetFileType(item.Name())
		if fileType == enums.IMAGE {
			imageFilePaths = append(imageFilePaths, filepath.Join(rootPath, item.Name()))
		}
	}

	return imageFilePaths
}

const webpExtension string = ".webp"

func isWebpImage(relativeImagePath string) bool {
	return filepath.Ext(relativeImagePath) == webpExtension
}

func hasLowQualityImageCounterpart(relativeImagePath string, allImagePaths []string) bool {
	if len(allImagePaths) == 0 {
		return false
	}

	possibleLowQualityImagePath := getLowQualityImagePath(relativeImagePath)
	return slices.Contains(allImagePaths, possibleLowQualityImagePath)
}

func isLowQualityImage(relativeImagePath string) bool {
	return strings.Contains(filepath.Base(relativeImagePath), fmt.Sprintf("%s%s", constants.LowQualityFileSuffix, webpExtension))
}

func getLowQualityImagePath(inputRelativeFilePath string) string {
	extension := filepath.Ext(inputRelativeFilePath)
	newFileSuffix := fmt.Sprintf("%s%s", constants.LowQualityFileSuffix, webpExtension)
	outputRelativeFilePath := strings.Replace(inputRelativeFilePath, extension, newFileSuffix, -1)
	return outputRelativeFilePath
}

func convertToWebp(inputRelativeFilePath string) (error, string) {

	conversionCommands := []Command{{command: "-quality", arg: "100"}, {command: "-define", arg: "webp:lossless=true"}}
	if isLossyImage(inputRelativeFilePath) {
		log.Printf("Image '%v' is lossy\n", inputRelativeFilePath)
		conversionCommands = []Command{{command: "-quality", arg: "95"}, {command: "-define", arg: "webp:method=6"}}
	} else {
		log.Printf("Image '%v' is lossless\n", inputRelativeFilePath)
	}
	stringifiedConversionCommands := convertImageMagickCommandsArrayToArgumentsArray(conversionCommands)

	extension := filepath.Ext(inputRelativeFilePath)
	outputRelativeFilePath := strings.Replace(inputRelativeFilePath, extension, webpExtension, -1)

	args := []string{inputRelativeFilePath}
	args = append(args, stringifiedConversionCommands...)
	args = append(args, outputRelativeFilePath)

	command := exec.Command("magick", args...)
	fmt.Printf("Executing: magick %s\n", strings.Join(args, " "))
	err := command.Run()
	if err != nil {
		log.Printf("Command failed: %v", err)
		log.Printf("Full command: magick %s", strings.Join(args, " "))
		return err, ""
	}

	return err, outputRelativeFilePath
}

func createLowQualityImage(inputRelativeFilePath string) (error, string) {
	conversionCommands := []Command{{command: "-resize", arg: "640x"}, {command: "-quality", arg: "10"}, {command: "-define", arg: "webp:method=6"}}
	stringifiedConversionCommands := convertImageMagickCommandsArrayToArgumentsArray(conversionCommands)

	outputRelativeFilePath := getLowQualityImagePath(inputRelativeFilePath)

	args := []string{inputRelativeFilePath}
	args = append(args, stringifiedConversionCommands...)
	args = append(args, outputRelativeFilePath)

	command := exec.Command("magick", args...)
	fmt.Printf("Executing: magick %s\n", strings.Join(args, " "))
	err := command.Run()
	if err != nil {
		log.Printf("Command failed: %v", err)
		log.Printf("Full command: magick %s", strings.Join(args, " "))
		return err, ""
	}

	return err, outputRelativeFilePath
}

// TODO MOVE TO UTILITEIS => MAYBE IMAGE UTILITIES
func isLossyImage(relativeImagePath string) bool {
	extension := filepath.Ext(relativeImagePath)
	return extension == ".jpg" || extension == ".jpeg"
}

func convertImageMagickCommandsArrayToArgumentsArray(commands []Command) []string {
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
