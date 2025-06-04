package imageResizer

import (
	models2 "backend/pkg/models"
	"testing"
)

//TODO

func Test_getResizeImageName(t *testing.T) {
	type GetResizeImageNameInput struct {
		name      string
		extension string
	}

	testData := []models2.TestData[GetResizeImageNameInput, string]{
		{Title: "Should return empty string on empty inputs", Input: GetResizeImageNameInput{name: "", extension: ""}, Expected: ""},
		{Title: "Should return empty string on empty name", Input: GetResizeImageNameInput{name: "", extension: ".jpg"}, Expected: ""},
		{Title: "Should return empty string on empty extension", Input: GetResizeImageNameInput{name: "file", extension: ""}, Expected: ""},
		{Title: "Should return filename with resize tag", Input: GetResizeImageNameInput{name: "file", extension: ".jpg"}, Expected: "file_resize.jpg"},
	}

	for _, data := range testData {

		t.Run(data.Title, func(t *testing.T) {
			// act
			mimeType := getResizeImageName(data.Input.name, data.Input.extension)

			// assert
			if mimeType != data.Expected {
				t.Errorf("Expected '%v' but received '%v'", data.Expected, mimeType)
			}
		})
	}
}
