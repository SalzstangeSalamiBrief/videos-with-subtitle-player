package lib

import (
	"backend/models"
	"fmt"
	"testing"
)

func TestGetFilenameWithoutExtension(t *testing.T) {
	inputs := []models.TestData[string, string]{
		{Input: "", Expected: ""},
		{Input: "example.mp3", Expected: "example"},
		{Input: "example.mp3.vtt", Expected: "example"},
		{Input: "example.wav", Expected: "example"},
		{Input: "example.wav.vtt", Expected: "example"},
		{Input: "example.mp4", Expected: "example"},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("input: '%v'", input.Input), func(t *testing.T) {
			result := GetFilenameWithoutExtension(input.Input)
			if result != input.Expected {
				t.Errorf("Expected '%v' but received '%v'", input.Expected, result)
			}
		})

	}
}
