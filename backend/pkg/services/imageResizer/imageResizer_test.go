package imageResizer

import (
	"backend/pkg/models"
	"path/filepath"
	"testing"
)

func Test_IsResizeFileName(t *testing.T) {
	testData := []models.TestData[string, bool]{
		{Title: "Should return empty string on empty inputs", Input: "", Expected: false},
		{Title: "Should return empty string on empty name", Input: "img", Expected: false},
		{Title: "Should return empty string on empty extension", Input: "C:\\myPath", Expected: false},
		{Title: "Should return empty string on empty extension", Input: "C:\\myPath\\image.png", Expected: false},
		{Title: "Should return empty string on empty extension", Input: "C:\\image_resize.png", Expected: true},
		{Title: "Should return empty string on empty extension", Input: "C:\\myPath\\image_resize.png", Expected: true},
	}

	for _, data := range testData {

		t.Run(data.Title, func(t *testing.T) {
			// act
			result := IsResizeFileName(data.Input)

			// assert
			if result != data.Expected {
				t.Errorf("Expected '%v' but received '%v'", data.Expected, result)
			}
		})
	}
}

func Test_getResizeImageName(t *testing.T) {
	type GetResizeImageNameInput struct {
		name      string
		extension string
	}

	testData := []models.TestData[GetResizeImageNameInput, string]{
		{Title: "Should return empty string on empty inputs", Input: GetResizeImageNameInput{name: "", extension: ""}, Expected: ""},
		{Title: "Should return empty string on empty name", Input: GetResizeImageNameInput{name: "", extension: ".jpg"}, Expected: ""},
		{Title: "Should return empty string on empty extension", Input: GetResizeImageNameInput{name: "file", extension: ""}, Expected: ""},
		{Title: "Should return filename with resize tag", Input: GetResizeImageNameInput{name: "file", extension: ".jpg"}, Expected: "file_resize.jpg"},
	}

	for _, data := range testData {

		t.Run(data.Title, func(t *testing.T) {
			// act
			result := getResizeImageName(data.Input.name, data.Input.extension)

			// assert
			if result != data.Expected {
				t.Errorf("Expected '%v' but received '%v'", data.Expected, result)
			}
		})
	}
}

func Test_addPathToResizeImage(t *testing.T) {
	type AddPathToResizeImageInput struct {
		inputFilePath       string
		resizeImageFileName string
	}

	testData := []models.TestData[AddPathToResizeImageInput, any]{
		{Title: "Should return empty string on empty inputs", Input: AddPathToResizeImageInput{inputFilePath: "", resizeImageFileName: ""}, Expected: ""},
		{Title: "Should return empty string on empty path", Input: AddPathToResizeImageInput{inputFilePath: "", resizeImageFileName: ".jpg"}, Expected: ""},
		{Title: "Should return empty string on empty name", Input: AddPathToResizeImageInput{inputFilePath: filepath.Join("a", "b", "c"), resizeImageFileName: ""}, Expected: ""},
		{Title: "Should return filename with resize tag", Input: AddPathToResizeImageInput{inputFilePath: filepath.Join("a", "b", "myfile.jgp"), resizeImageFileName: "file.jpg"}, Expected: filepath.Join("a", "b", "file.jpg")},
	}

	for _, data := range testData {

		t.Run(data.Title, func(t *testing.T) {
			// act
			result := addPathToResizeImage(data.Input.inputFilePath, data.Input.resizeImageFileName)

			// assert
			if result != data.Expected {
				t.Errorf("Expected '%v' but received '%v'", data.Expected, result)
			}
		})
	}
}

// TODO Test_GetFilenameAndExtensionParts
// TODO Test_GetResizeImagePath
