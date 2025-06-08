package imageResizer

import (
	"backend/pkg/models"
	"path/filepath"
	"testing"
)

func Test_IsResizeFileName(t *testing.T) {
testData := []models.TestData[string, bool]{
		{Title: "Should return false for empty input", Input: "", Expected: false},
		{Title: "Should return false for filename without resize suffix", Input: "img", Expected: false},
		{Title: "Should return false for path without filename", Input: "C:\\myPath", Expected: false},
		{Title: "Should return false for normal image filename", Input: "C:\\myPath\\image.png", Expected: false},
		{Title: "Should return true for resized image filename", Input: "C:\\image_resize.png", Expected: true},
		{Title: "Should return true for resized image with path", Input: "C:\\myPath\\image_resize.png", Expected: true},
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

func Test_getFilenameAndExtensionParts(t *testing.T) {
testData := []models.TestData[string, [2]string]{
 		{Title: "Should return empty name and extension for empty path", Input: "", Expected: [2]string{"", ""}},
		{Title: "Should return name and empty extension for path without extension", Input: filepath.Join("myPath", "image"), Expected: [2]string{"image", ""}},
		{Title: "Should return name and extension for normal file path", Input: filepath.Join("myPath", "image.png"), Expected: [2]string{"image", ".png"}},
		{Title: "Should handle path with multiple dots", Input: filepath.Join("myPath", "my.image.png"), Expected: [2]string{"my.image", ".png"}},
		{Title: "Should return correct name and extension with nested folders", Input: filepath.Join("folder", "subfolder", "file.ext"), Expected: [2]string{"file", ".ext"}},
 	}

	for _, data := range testData {
		t.Run(data.Title, func(t *testing.T) {
			// act
			name, ext := getFilenameAndExtensionParts(data.Input)

			// assert
			if name != data.Expected[0] || ext != data.Expected[1] {
				t.Errorf("Expected name='%v' and ext='%v', but got name='%v' and ext='%v'",
					data.Expected[0], data.Expected[1], name, ext)
			}
		})
	}
}

func Test_getResizeImagePath(t *testing.T) {
testData := []models.TestData[string, string]{
 		{
 			Title:    "Should return resized path with '_resize' suffix",
			Input:    filepath.Join("images", "image.png"),
			Expected: filepath.Join("images", "image_resize.png"),
 		},
 		{
 			Title:    "Should handle files with multiple dots",
			Input:    filepath.Join("images", "my.image.png"),
			Expected: filepath.Join("images", "my.image_resize.png"),
 		},
 		{
 			Title:    "Should handle file without extension",
			Input:    filepath.Join("images", "image"),
 			Expected: "",
 		},
 	}

	for _, data := range testData {
		t.Run(data.Title, func(t *testing.T) {
			// Act
			result := getResizeImagePath(data.Input)

			// Assert
			if result != data.Expected {
				t.Errorf("Expected '%v', got '%v'", data.Expected, result)
			}
		})
	}
}
