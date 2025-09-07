package webpTransformer

import (
	"log"
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

func NewWebpTransformerConfiguration() *WebpTransformerConfiguration {
	return &WebpTransformerConfiguration{
		rootPath:                   loadRootPath(),
		shouldDeleteNonWebpImages:  loadShouldDeleteNonWebpImages(),
		executionIntervalInMinutes: loadExecutionIntervalInMinutes(),
	}
}

func loadRootPath() string {
	rp := os.Getenv(rootPathKey)
	if rp == "" {
		log.Fatalf("Could not load env variable '%v'", rootPathKey)
	}

	return rp
}

func loadShouldDeleteNonWebpImages() bool {
	stringifiedValue := os.Getenv(shouldDeleteNonWebpImagesKey)
	if stringifiedValue == "" {
		log.Fatalf("Could not load env variable '%v'", shouldDeleteNonWebpImagesKey)
	}

	parsedValue, err := strconv.ParseBool(stringifiedValue)
	if err != nil {
		log.Fatalf("Could not parse '%v'", stringifiedValue)
	}

	return parsedValue
}

func loadExecutionIntervalInMinutes() int64 {
	stringifiedValue := os.Getenv(executionIntervalInMinutesKey)
	if stringifiedValue == "" {
		log.Fatalf("Could not load env variable '%v'", executionIntervalInMinutesKey)
	}

	parsedValue, err := strconv.ParseInt(stringifiedValue, 10, 64)
	if err != nil {
		log.Fatalf("Could not parse '%v'", stringifiedValue)
	}

	return parsedValue
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
