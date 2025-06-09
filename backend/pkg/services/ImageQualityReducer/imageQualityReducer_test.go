package ImageQualityReducer

import (
	"backend/pkg/models"
	"fmt"
	"path/filepath"
	"reflect"
	"testing"
)

func Test_IsResizeFileName(t *testing.T) {
	testData := []models.TestData[string, bool]{
		{Title: "Should return false for empty input", Input: "", Expected: false},
		{Title: "Should return false for filename without resize suffix", Input: "img", Expected: false},
		{Title: "Should return false for path without filename", Input: filepath.Join("C:", "myPath"), Expected: false},
		{Title: "Should return false for normal image filename", Input: filepath.Join("C:", "myPath", "image.png"), Expected: false},
		{Title: "Should return true for resized image filename", Input: filepath.Join("C:", fmt.Sprintf("image%v.png", lowQualityFileSuffix)), Expected: true},
		{Title: "Should return true for resized image with path", Input: filepath.Join("C:", "myPath", fmt.Sprintf("image%v.png", lowQualityFileSuffix)), Expected: true},
	}

	for _, data := range testData {

		t.Run(data.Title, func(t *testing.T) {
			// act
			result := IsLowQualityFileName(data.Input)

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
		{Title: "Should return filename with resize tag", Input: GetResizeImageNameInput{name: "file", extension: ".jpg"}, Expected: fmt.Sprintf("file%v.jpg", lowQualityFileSuffix)},
	}

	for _, data := range testData {

		t.Run(data.Title, func(t *testing.T) {
			// act
			result := getLowQualityImageName(data.Input.name, data.Input.extension)

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
			result := addPathToLowQualityImage(data.Input.inputFilePath, data.Input.resizeImageFileName)

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
			result := getLowQualityImagePath(data.Input)

			// Assert
			if result != data.Expected {
				t.Errorf("Expected '%v', got '%v'", data.Expected, result)
			}
		})
	}
}

func Test_convertImageMagickCommandsArrayToArgumentsArray(t *testing.T) {
	testData := []models.TestData[[]ImageMagickCommand, []string]{
		{
			Title:    "Should return empty slice for empty input",
			Input:    []ImageMagickCommand{},
			Expected: []string{},
		},
		{
			Title: "Should convert single command-arg pair",
			Input: []ImageMagickCommand{
				{command: "-resize", arg: "100x100"},
			},
			Expected: []string{"-resize", "100x100"},
		},
		{
			Title: "Should convert multiple command-arg pairs",
			Input: []ImageMagickCommand{
				{command: "-resize", arg: "200x200"},
				{command: "-quality", arg: "85"},
				{command: "-strip", arg: "true"},
			},
			Expected: []string{"-resize", "200x200", "-quality", "85", "-strip", "true"},
		},
		{
			Title: "Should skip empty command and arg",
			Input: []ImageMagickCommand{
				{command: "", arg: ""},
				{command: "-resize", arg: ""},
				{command: "", arg: "85"},
			},
			Expected: []string{"-resize", "85"},
		},
		{
			Title: "Should include only command if arg is empty",
			Input: []ImageMagickCommand{
				{command: "-resize", arg: ""},
			},
			Expected: []string{"-resize"},
		},
		{
			Title: "Should include only arg if command is empty",
			Input: []ImageMagickCommand{
				{command: "", arg: "85"},
			},
			Expected: []string{"85"},
		},
	}

	for _, data := range testData {
		t.Run(data.Title, func(t *testing.T) {
			// act
			result := convertImageMagickCommandsArrayToArgumentsArray(data.Input)

			// assert
			if !reflect.DeepEqual(result, data.Expected) {
				t.Errorf("Expected %v, but got %v", data.Expected, result)
			}
		})
	}
}
