package utilities

import (
	"backend/models"
	"testing"
)

func TestGetContentTypeHeaderMimeType(t *testing.T) {
	testData := [][2]string{
		{"path/to/file.vtt", "text/vtt"},
		{"path/to/file.mp4", "video/mp4"},
		{"path/to/file.mp3", "audio/mpeg"},
		{"file", ""},
	}

	for _, data := range testData {

		t.Run(data[1], func(t *testing.T) {
			// arrange
			file := models.FileTreeItem{
				Id:   "1",
				Path: data[0],
			}

			// act
			mimeType := GetContentTypeHeaderMimeType(file)

			// assert
			if mimeType != data[1] {
				t.Errorf("Expected '%v' but received '%v'", data[1], mimeType)
			}
		})
	}
}
