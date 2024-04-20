package usecases

import (
	"backend/models"
	"testing"
)

func Test_GetContentTypeHeaderMimeType(t *testing.T) {
	testData := []models.TestData[string, string]{
		{Input: "path/to/file.vtt", Expected: "text/vtt"},
		{Input: "path/to/file.mp4", Expected: "video/mp4"},
		{Input: "path/to/file.mp3", Expected: "audio/mpeg"},
		{Input: "file", Expected: ""},
	}

	for _, data := range testData {

		t.Run(data.Input, func(t *testing.T) {
			// arrange
			file := models.FileTreeItem{
				Id:   "1",
				Path: data.Input,
			}

			// act
			mimeType := GetContentTypeHeaderMimeType(file)

			// assert
			if mimeType != data.Expected {
				t.Errorf("Expected '%v' but received '%v'", data.Expected, mimeType)
			}
		})
	}
}
