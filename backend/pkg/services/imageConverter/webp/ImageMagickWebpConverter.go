package webp

import (
	"backend/pkg/enums"
	"backend/pkg/services/imageConverter/constants"
	"backend/pkg/services/imageConverter/models"
	"backend/pkg/utilities"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

type ExecuteWebpConversionConfiguration struct {
	RootPath string
}

func ExecuteWebpConversion(configuration ExecuteWebpConversionConfiguration) error {
	err, allImages := traverseFileTreeToGetImages(configuration.RootPath)
	if err != nil {
		return err
	}

	for _, image := range allImages {
		if isLowQualityImage(image) {
			log.Printf("Image %s is low quality. Skip processing \n", image)
			continue
		}

		if !isWebpImage(image) {
			webpConversionError, _ := convertToWebp(image)
			// TODO REMOVE SOURCE FILE?

			if webpConversionError != nil {
				log.Fatal(webpConversionError)
			}
		}

		if !hasLowQualityImage(image, allImages) {
			log.Printf("Image %s has no low quality counterpart. Create low quality image \n", image)
			createLowQualityImageError, _ := createLowQualityImage(image)
			if createLowQualityImageError != nil {
				log.Fatal(createLowQualityImageError)
			}
		}

	}

	log.Printf("Finish webp conversion. Converted '%v' files\n", len(allImages))
	return nil
}

func isWebpImage(relativeImagePath string) bool {
	return filepath.Ext(relativeImagePath) == constants.WebpExtension
}

func hasLowQualityImage(relativeImagePath string, allImagePaths []string) bool {
	if len(allImagePaths) == 0 {
		return false
	}

	possibleLowQualityImagePath := getLowQualityImagePath(relativeImagePath)
	return slices.Contains(allImagePaths, possibleLowQualityImagePath)
}

func isLowQualityImage(relativeImagePath string) bool {
	return strings.Contains(filepath.Base(relativeImagePath), fmt.Sprintf("%s%s", constants.LowQualityFileSuffix, constants.WebpExtension))
}

// TODO MOVE TO UTILITEIS => MAYBE IMAGE UTILITIES
func isLossyImage(relativeImagePath string) bool {
	extension := filepath.Ext(relativeImagePath)
	return extension == ".jpg" || extension == ".jpeg"
}

func getLowQualityImagePath(inputRelativeFilePath string) string {
	extension := filepath.Ext(inputRelativeFilePath)
	newFileSuffix := fmt.Sprintf("%s%s", constants.LowQualityFileSuffix, constants.WebpExtension)
	outputRelativeFilePath := strings.Replace(inputRelativeFilePath, extension, newFileSuffix, -1)
	return outputRelativeFilePath
}

func convertToWebp(inputRelativeFilePath string) (error, string) {

	conversionCommands := []models.ImageCLICommand{{Command: "-quality", Arg: "100"}, {Command: "-define", Arg: "webp:lossless=true"}}
	if isLossyImage(inputRelativeFilePath) {
		log.Printf("Image '%v' is lossy\n", inputRelativeFilePath)
		conversionCommands = []models.ImageCLICommand{{Command: "-quality", Arg: "95"}, {Command: "-define", Arg: "webp:method=6"}}
	} else {
		log.Printf("Image '%v' is lossless\n", inputRelativeFilePath)
	}
	stringifiedConversionCommands := models.ConvertImageMagickCommandsArrayToArgumentsArray(conversionCommands)

	extension := filepath.Ext(inputRelativeFilePath)
	outputRelativeFilePath := strings.Replace(inputRelativeFilePath, extension, constants.WebpExtension, -1)

	args := []string{inputRelativeFilePath}
	args = append(args, stringifiedConversionCommands...)
	args = append(args, outputRelativeFilePath)

	command := exec.Command("magick", args...)
	fmt.Printf("Executing: magick %s\n", strings.Join(args, " "))
	err := command.Run()
	if err != nil {
		log.Printf("Command failed: %v", err)
		log.Printf("Full Command: magick %s", strings.Join(args, " "))
		return err, ""
	}

	return err, outputRelativeFilePath
}

func createLowQualityImage(inputRelativeFilePath string) (error, string) {
	conversionCommands := []models.ImageCLICommand{{Command: "-resize", Arg: "640x"}, {Command: "-quality", Arg: "10"}, {Command: "-define", Arg: "webp:method=6"}}
	stringifiedConversionCommands := models.ConvertImageMagickCommandsArrayToArgumentsArray(conversionCommands)

	outputRelativeFilePath := getLowQualityImagePath(inputRelativeFilePath)

	args := []string{inputRelativeFilePath}
	args = append(args, stringifiedConversionCommands...)
	args = append(args, outputRelativeFilePath)

	command := exec.Command("magick", args...)
	fmt.Printf("Executing: magick %s\n", strings.Join(args, " "))
	err := command.Run()
	if err != nil {
		log.Printf("Command failed: %v", err)
		log.Printf("Full Command: magick %s", strings.Join(args, " "))
		return err, ""
	}

	return err, outputRelativeFilePath
}

func traverseFileTreeToGetImages(rootPath string) (error, []string) {
	imageFilePaths := make([]string, 0)

	content, readDirError := os.ReadDir(rootPath)
	if readDirError != nil {
		return readDirError, []string{}
	}

	for _, item := range content {
		fullPath := filepath.Join(rootPath, item.Name())
		if item.IsDir() {
			subError, subContent := traverseFileTreeToGetImages(fullPath)
			if subError != nil {
				return subError, []string{}
			}

			if len(subContent) > 0 {
				imageFilePaths = append(imageFilePaths, subContent...)
			}
		}

		fileType := utilities.GetFileType(item.Name())
		if fileType == enums.IMAGE {
			imageFilePaths = append(imageFilePaths, fullPath)
		}
	}

	return nil, imageFilePaths
}
