package configuration

import "backend/internal/configuration/utilities"

var shouldDeleteNonWebpImagesDefault = false

type WebpTransformerConfiguration struct {
	rootPath                   string
	shouldDeleteNonWebpImages  bool
	executionIntervalInMinutes int64
}

func NewWebpTransformerConfiguration() (*WebpTransformerConfiguration, error) {
	rootPath, err := utilities.GetEnvironmentString("RootPath", true, nil)
	if err != nil {
		return nil, err
	}

	executionIntervalInMinutes, err := utilities.GetEnvironmentInt("IntervalInMinutes", true, nil)
	if err != nil {
		return nil, err
	}

	shouldDeleteNonWebpImages, err := utilities.GetEnvironmentBoolean("ShouldDeleteNonWebpImages", false, &shouldDeleteNonWebpImagesDefault)
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
