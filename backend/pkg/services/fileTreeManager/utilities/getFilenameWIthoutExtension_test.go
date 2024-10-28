package utilities

import (
	"backend/pkg/models"
	"testing"
)

func Test_GetFilenameWithoutExtension(t *testing.T) {
	inputs := []models.TestData[string, string]{
		{Title: "Should return empty string on empty input", Input: "", Expected: ""},
		{Title: "Should return file name of mp3 input", Input: "example.mp3", Expected: "example"},
		{Title: "Should return file name of mp3 subtitle input", Input: "example.mp3.vtt", Expected: "example"},
		{Title: "Should return file name of wav input", Input: "example.wav", Expected: "example"},
		{Title: "Should return file name of wav subtitle input", Input: "example.wav.vtt", Expected: "example"},
		{Title: "Should return file name of mp4 file", Input: "example.mp4", Expected: "example"},
		{Title: "Should return file name that contains one dot for a media file", Input: "1. example.mp4", Expected: "1. example"},
		{Title: "Should return file name that contains one dot for a subtitle file", Input: "1. example.wav.vtt", Expected: "1. example"},
		{Title: "Should return file name that contains one dot for a media file", Input: "1. 2. 3. example.mp4", Expected: "1. 2. 3. example"},
		{Title: "Should return file name that contains multiple dots for a subtitle file", Input: "1. 2. 3. example.wav.vtt", Expected: "1. 2. 3. example"},
	}

	for _, input := range inputs {
		t.Run(input.Title, func(t *testing.T) {
			result := GetFilenameWithoutExtension(input.Input)
			if result != input.Expected {
				t.Errorf("Expected '%v' but received '%v'", input.Expected, result)
			}
		})

	}
}
