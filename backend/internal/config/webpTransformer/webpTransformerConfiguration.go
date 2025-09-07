package webpTransformer

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type WebpTransformerConfiguration struct {
	rootPath                   string
	shouldDeleteNonWebpImages  bool
	executionIntervalInMinutes int64
}

const rootPathKey string = "RootPath"
const shouldDeleteNonWebpImagesKey string = "ShouldDeleteNonWebpImages"
const executionIntervalInMinutesKey string = "intervalInMinutes"

func NewWebpTransformerConfiguration() (*WebpTransformerConfiguration, error) {
	rootPath, err := loadRootPath()
	if err != nil {
		return nil, err
	}

	shouldDeleteNonWebpImages, err := loadShouldDeleteNonWebpImages()
	if err != nil {
		return nil, err
	}

	executionIntervalInMinutes, err := loadExecutionIntervalInMinutes()
	if err != nil {
		return nil, err
	}

	return &WebpTransformerConfiguration{
		rootPath:                   rootPath,
		shouldDeleteNonWebpImages:  shouldDeleteNonWebpImages,
		executionIntervalInMinutes: executionIntervalInMinutes,
	}, nil
}

func loadRootPath() (string, error) {
	rp := os.Getenv(rootPathKey)
	if rp == "" {
		return "", errors.New(fmt.Sprintf("Could not load env variable '%v'", rootPathKey))
	}

	return rp, nil
}

func loadShouldDeleteNonWebpImages() (bool, error) {
	stringifiedValue := os.Getenv(shouldDeleteNonWebpImagesKey)
	if stringifiedValue == "" {
		return false, errors.New(fmt.Sprintf("Could not load env variable '%v'", shouldDeleteNonWebpImagesKey))
	}

	parsedValue, err := strconv.ParseBool(stringifiedValue)
	if err != nil {
		return false, errors.New(fmt.Sprintf("Could not parse '%v'", stringifiedValue))
	}

	return parsedValue, nil
}

func loadExecutionIntervalInMinutes() (int64, error) {
	stringifiedValue := os.Getenv(executionIntervalInMinutesKey)
	if stringifiedValue == "" {
		return 0, errors.New(fmt.Sprintf("Could not load env variable '%v'", executionIntervalInMinutesKey))
	}

	parsedValue, err := strconv.ParseInt(stringifiedValue, 10, 64)
	if err != nil {
		return 0, err
	}

	if parsedValue < 1 {
		return 0, errors.New(fmt.Sprintf("The value of '%v' '%v' cannot be lesser than '1'", executionIntervalInMinutesKey, parsedValue))
	}

	return parsedValue, nil
}

func (configuration *WebpTransformerConfiguration) GetRootPath() string {
	return configuration.rootPath
}

func (configuration *WebpTransformerConfiguration) GetShouldDeleteNonWebpImages() bool {
	return configuration.shouldDeleteNonWebpImages
}

func (configuration *WebpTransformerConfiguration) GetExecutionIntervalInMinutes() int64 {
	return configuration.executionIntervalInMinutes
}
