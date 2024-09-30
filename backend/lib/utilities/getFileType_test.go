package lib

import (
	"backend/enums"
	"backend/models"
	"testing"
)

func Test_GetFileType(t *testing.T) {
	inputs := []models.TestData[string, enums.FileType]{
		{Title: "Return unknown on empty input", Input: "", Expected: enums.UNKNOWN},
		{Title: "Return unknown for text file", Input: "example.txt", Expected: enums.UNKNOWN},
		{Title: "Return subtitle for vtt file", Input: "example.vtt", Expected: enums.SUBTITLE},
		{Title: "Return audio for mp3 file", Input: "example.mp3", Expected: enums.AUDIO},
		{Title: "Return audio for wav file", Input: "example.wav", Expected: enums.AUDIO},
		{Title: "Return video for mp4 file", Input: "example.mp4", Expected: enums.VIDEO},
		{Title: "Return unknown for svg file", Input: "example.svg", Expected: enums.UNKNOWN},
		{Title: "Return video for jpg file", Input: "example.jpg", Expected: enums.IMAGE},
		{Title: "Return video for jpeg file", Input: "example.jpeg", Expected: enums.IMAGE},
		{Title: "Return video for avif file", Input: "example.avif", Expected: enums.IMAGE},
		{Title: "Return video for webp file", Input: "example.webp", Expected: enums.IMAGE},
		{Title: "Return video for png file", Input: "example.png", Expected: enums.IMAGE},
	}

	for _, input := range inputs {
		t.Run(input.Title, func(t *testing.T) {
			result := GetFileType(input.Input)

			if result != input.Expected {
				t.Errorf("Expected '%v' but received '%v'", input.Expected, result)
			}
		})
	}
}
