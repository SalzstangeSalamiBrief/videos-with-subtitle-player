package webpTransformer

import (
	"backend/internal/configuration"
)

var shouldDeleteNonWebpImagesDefault = false

type WebpTransformerConfiguration struct {
	rootPath                   string
	shouldDeleteNonWebpImages  bool
	executionIntervalInMinutes int64
}

func NewWebpTransformerConfiguration() (*WebpTransformerConfiguration, error) {
	rootPath, err := configuration.GetEnvironmentString("RootPath", true, nil)
	if err != nil {
		return nil, err
	}

	executionIntervalInMinutes, err := configuration.GetEnvironmentInt("IntervalInMinutes", true, nil)
	if err != nil {
		return nil, err
	}

	shouldDeleteNonWebpImages, err := configuration.GetEnvironmentBoolean("ShouldDeleteNonWebpImages", false, &shouldDeleteNonWebpImagesDefault)
	if err != nil {
		return nil, err
	}

	return &WebpTransformerConfiguration{
		rootPath,
		shouldDeleteNonWebpImages,
		executionIntervalInMinutes,
	}, nil
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
