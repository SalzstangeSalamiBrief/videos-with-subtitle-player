package utilities

import (
	"backend/models"
	"fmt"
	"testing"
)

func Test_getStringifiedRange(t *testing.T) {
	inputs := []models.TestData[string, [2]string]{
		{Input: "", Expected: [2]string{"", ""}},
		{Input: "bytes=", Expected: [2]string{"", ""}},
		{Input: "bytes=1", Expected: [2]string{"1", ""}},
		{Input: "bytes=1-2", Expected: [2]string{"1", "2"}},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("input: %v", input.Input), func(t *testing.T) {
			stringifiedStart, stringifiedEnd := getStringifiedRange(input.Input)
			if stringifiedStart != input.Expected[0] {
				t.Errorf("Expected '%v' in as start value but received '%v'", input.Expected[0], stringifiedStart)
			}

			if stringifiedEnd != input.Expected[1] {
				t.Errorf("Expected '%v' in as end value but received '%v'", input.Expected[1], stringifiedStart)
			}
		})
	}
}

func Test_getStart(t *testing.T) {
	inputs := []models.TestData[string, models.ValueErrorCombination[int64]]{
		{Input: "", Expected: models.ValueErrorCombination[int64]{Value: 0, Error: nil}},
		{Input: "abc", Expected: models.ValueErrorCombination[int64]{Value: 0, Error: fmt.Errorf("")}},
		{Input: "-1", Expected: models.ValueErrorCombination[int64]{Value: -1, Error: nil}},
		{Input: "1", Expected: models.ValueErrorCombination[int64]{Value: 1, Error: nil}},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("input '%v'", input.Input), func(t *testing.T) {
			resultValue, resultError := getStart(input.Input)
			if resultValue != input.Expected.Value {
				t.Errorf("Expected '%v' but received '%v'", input.Expected.Value, resultValue)
			}

			if input.Expected.Error != nil && resultError == nil {
				t.Error("Expected an error but received none")
			}

			if input.Expected.Error == nil && resultError != nil {
				t.Errorf("Expected an error but received none")
			}
		})
	}
}

func Test_getEnd(t *testing.T) {
	inputs := []models.TestData[getEndInput, models.ValueErrorCombination[int64]]{
		{Input: getEndInput{stringifiedEnd: "", fileSize: 1024, chunkSize: 1024, start: 0}, Expected: models.ValueErrorCombination[int64]{Value: 1024, Error: nil}},
		{Input: getEndInput{stringifiedEnd: "0", fileSize: 1024, chunkSize: 1024, start: 0}, Expected: models.ValueErrorCombination[int64]{Value: 1024, Error: nil}},
		{Input: getEndInput{stringifiedEnd: "a", fileSize: 1024, chunkSize: 1024, start: 0}, Expected: models.ValueErrorCombination[int64]{Value: 0, Error: fmt.Errorf("")}},
		{Input: getEndInput{stringifiedEnd: "10", fileSize: 1024, chunkSize: 1024, start: 0}, Expected: models.ValueErrorCombination[int64]{Value: 10, Error: nil}},
		{Input: getEndInput{stringifiedEnd: "2048", fileSize: 1024, chunkSize: 1024, start: 0}, Expected: models.ValueErrorCombination[int64]{Value: 1024, Error: nil}},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("input '%#v'", input.Input), func(t *testing.T) {
			resultValue, resultError := getEnd(input.Input)
			if resultValue != input.Expected.Value {
				t.Errorf("Expected '%v' but received '%v'", input.Expected.Value, resultValue)
			}

			if input.Expected.Error != nil && resultError == nil {
				t.Error("Expected an error but received none")
			}

			if input.Expected.Error == nil && resultError != nil {
				t.Errorf("Expected an error but received none")
			}
		})
	}
}
