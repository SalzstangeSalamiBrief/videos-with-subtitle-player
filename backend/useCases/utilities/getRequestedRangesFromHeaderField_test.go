package usecases

import (
	"backend/models"
	"testing"
)

func Test_getStringifiedRange(t *testing.T) {
	inputs := []models.TestData[string, [2]string]{
		{Title: "Should return empty tuple based on empty input", Input: "", Expected: [2]string{"", ""}},
		{Title: "Should return empty tuple based on invalid range string", Input: "bytes=", Expected: [2]string{"", ""}},
		{Title: "Should return tuple with a start", Input: "bytes=1", Expected: [2]string{"1", ""}},
		{Title: "Should return a tuple with a start and end", Input: "bytes=1-2", Expected: [2]string{"1", "2"}},
	}

	for _, input := range inputs {
		t.Run(input.Title, func(t *testing.T) {
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
		{Title: "Should return 0 and no error empty input", Input: "", Expected: models.ValueErrorCombination[int64]{Value: 0, HasError: false}},
		{Title: "Should return 0 and an error on invalid stringified int", Input: "abc", Expected: models.ValueErrorCombination[int64]{Value: 0, HasError: true}},
		{Title: "Should return -1 on valid negative input", Input: "-1", Expected: models.ValueErrorCombination[int64]{Value: -1, HasError: false}},
		{Title: "Should return 1 on valid positive input", Input: "1", Expected: models.ValueErrorCombination[int64]{Value: 1, HasError: false}},
	}

	for _, input := range inputs {
		t.Run(input.Title, func(t *testing.T) {
			resultValue, resultError := getStart(input.Input)
			if resultValue != input.Expected.Value {
				t.Errorf("Expected '%v' but received '%v'", input.Expected.Value, resultValue)
			}

			if input.Expected.HasError && resultError == nil {
				t.Error("Expected an error but received none")
			}

			if !input.Expected.HasError && resultError != nil {
				t.Errorf("Expected an error but received none")
			}
		})
	}
}

func Test_getEnd(t *testing.T) {
	inputs := []models.TestData[getEndInput, models.ValueErrorCombination[int64]]{
		{Title: "Should return file size on empty end", Input: getEndInput{stringifiedEnd: "", fileSize: 1024, chunkSize: 1024, start: 0}, Expected: models.ValueErrorCombination[int64]{Value: 1024, HasError: false}},
		{Title: "Should return file size on end string equal to zero", Input: getEndInput{stringifiedEnd: "0", fileSize: 1024, chunkSize: 1024, start: 0}, Expected: models.ValueErrorCombination[int64]{Value: 1024, HasError: false}},
		{Title: "Should return 0 and an error on invalid end string", Input: getEndInput{stringifiedEnd: "a", fileSize: 1024, chunkSize: 1024, start: 0}, Expected: models.ValueErrorCombination[int64]{Value: 0, HasError: true}},
		{Title: "Should return the end specified in the input", Input: getEndInput{stringifiedEnd: "10", fileSize: 1024, chunkSize: 1024, start: 0}, Expected: models.ValueErrorCombination[int64]{Value: 10, HasError: false}},
		{Title: "Should return the chunk size on end greater than the file size", Input: getEndInput{stringifiedEnd: "2048", fileSize: 1024, chunkSize: 1024, start: 0}, Expected: models.ValueErrorCombination[int64]{Value: 1024, HasError: false}},
	}

	for _, input := range inputs {
		t.Run(input.Title, func(t *testing.T) {
			resultValue, resultError := getEnd(input.Input)
			if resultValue != input.Expected.Value {
				t.Errorf("Expected '%v' but received '%v'", input.Expected.Value, resultValue)
			}

			if input.Expected.HasError && resultError == nil {
				t.Error("Expected an error but received none")
			}

			if !input.Expected.HasError && resultError != nil {
				t.Errorf("Expected an error but received none")
			}
		})
	}
}

func Test_GetRequestedRangesFromHeaderField(t *testing.T) {
	inputs := []models.TestData[GetRequestRangesInput, [2]int64]{
		{Title: "Should return zero to chunk size on empty range header", Input: GetRequestRangesInput{"", 1024, 0}, Expected: [2]int64{0, 1024}},
		{Title: "Should return zero to chuck size on incomplete range header", Input: GetRequestRangesInput{"bytes=", 1024, 0}, Expected: [2]int64{0, 1024}},
		{Title: "Should return start and sum of chunk size with start on range header with only start", Input: GetRequestRangesInput{"bytes=1-", 1024, 0}, Expected: [2]int64{1, 1025}},
		{Title: "Should return start and file size on header greater than chunk size", Input: GetRequestRangesInput{"bytes=1-2000", 1024, 512}, Expected: [2]int64{1, 512}},
		{Title: "Should return start and end from range header", Input: GetRequestRangesInput{"bytes=1-2000", 1024, 16000}, Expected: [2]int64{1, 2000}},
	}

	for _, input := range inputs {
		t.Run(input.Title, func(t *testing.T) {
			start, end := GetRequestedRangesFromHeaderField(input.Input)
			if start != input.Expected[0] {
				t.Errorf("Expected '%v' but received '%v'", input.Expected[0], start)
			}

			if end != input.Expected[1] {
				t.Errorf("Expected '%v' but received '%v'", input.Expected[1], end)
			}
		})
	}
}
