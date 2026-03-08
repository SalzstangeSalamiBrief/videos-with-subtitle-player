package utilities

import (
	"backend/pkg/enums/fileType"
	"backend/pkg/models"
	"testing"
)

func Test_GetFileType(t *testing.T) {
	inputs := []models.TestData[string, fileType.FileType]{
		{Title: "Return unknown on empty input", Input: "", Expected: fileType.UNKNOWN},
		{Title: "Return unknown for text file", Input: "example.txt", Expected: fileType.UNKNOWN},
		{Title: "Return subtitle for vtt file", Input: "example.vtt", Expected: fileType.SUBTITLE},
		{Title: "Return audio for mp3 file", Input: "example.mp3", Expected: fileType.AUDIO},
		{Title: "Return audio for wav file", Input: "example.wav", Expected: fileType.AUDIO},
		{Title: "Return video for mp4 file", Input: "example.mp4", Expected: fileType.VIDEO},
		{Title: "Return unknown for svg file", Input: "example.svg", Expected: fileType.UNKNOWN},
		{Title: "Return video for jpg file", Input: "example.jpg", Expected: fileType.IMAGE},
		{Title: "Return video for jpeg file", Input: "example.jpeg", Expected: fileType.IMAGE},
		{Title: "Return video for avif file", Input: "example.avif", Expected: fileType.IMAGE},
		{Title: "Return video for webp file", Input: "example.webp", Expected: fileType.IMAGE},
		{Title: "Return video for png file", Input: "example.png", Expected: fileType.IMAGE},
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
