package usecases

import (
	"backend/models"
	"testing"
)

func Test_GetContentTypeHeaderMimeType(t *testing.T) {
	// TODO ADD TITLE TO TESTDATA => ADD TITLE TO TEST DATA LIKE "SHOULD ..."
	testData := []models.TestData[string, string]{
		{Title: "Should return subtitle mime type", Input: "path/to/file.vtt", Expected: "text/vtt"},
		{Title: "Should return mp4 mime type", Input: "path/to/file.mp4", Expected: "video/mp4"},
		{Title: "Should return mpeg mime type", Input: "path/to/file.mp3", Expected: "audio/mpeg"},
		{Title: "Should return empty string", Input: "file", Expected: ""},
	}

	for _, data := range testData {

		t.Run(data.Title, func(t *testing.T) {
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
