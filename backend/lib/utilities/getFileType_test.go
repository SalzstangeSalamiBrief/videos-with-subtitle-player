package lib

import (
	"backend/enums"
	"backend/models"
	"fmt"
	"testing"
)

func Test_getFileType(t *testing.T) {
	inputs := []models.TestData[string, enums.FileType]{
		{Input: "", Expected: enums.UNKNOWN},
		{Input: "example.txt", Expected: enums.UNKNOWN},
		{Input: "example.vtt", Expected: enums.SUBTITLE},
		{Input: "example.mp3", Expected: enums.AUDIO},
		{Input: "example.wav", Expected: enums.AUDIO},
		{Input: "example.mp4", Expected: enums.VIDEO},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("input: '%v'", input.Input), func(t *testing.T) {
			result := GetFileType(input.Input)

			if result != input.Expected {
				t.Errorf("Expected '%v' but received '%v'", input.Expected, result)
			}
		})
	}
}
