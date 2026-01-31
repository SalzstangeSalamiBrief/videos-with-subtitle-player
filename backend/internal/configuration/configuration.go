package configuration

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetEnvironmentString(key string, isRequired bool, defaultValue *string) (string, error) {
	value := strings.Trim(os.Getenv(key), " ")

	if value == "" {
		if isRequired {
			return "", errors.New(fmt.Sprintf("Could not load required string environment variable '%v'", key))

		}

		if defaultValue != nil && *defaultValue != "" {
			return *defaultValue, nil
		}
	}

	return value, nil
}

func GetEnvironmentInt(key string, isRequired bool, defaultValue *int64) (int64, error) {
	valueString := strings.Trim(os.Getenv(key), " ")
	if valueString == "" {
		if isRequired {
			return 0, errors.New(fmt.Sprintf("Could not load required int environment variable '%v'", key))
		}

		if defaultValue != nil {
			return *defaultValue, nil
		}
	}

	valueInt, convertError := strconv.ParseInt(valueString, 10, 64)
	return valueInt, convertError
}

func GetEnvironmentBoolean(key string, isRequired bool, defaultValue *bool) (bool, error) {
	valueString := strings.Trim(os.Getenv(key), " ")
	if valueString == "" {
		if isRequired {
			return false, errors.New(fmt.Sprintf("Could not load required int environment variable '%v'", key))
		}

		if defaultValue != nil {
			return *defaultValue, nil
		}
	}

	valueBool, convertError := strconv.ParseBool(valueString)
	return valueBool, convertError
}
