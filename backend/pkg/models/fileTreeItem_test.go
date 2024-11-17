package models

import (
	"path/filepath"
	"testing"
)

func Test_GetPartsOfPath(t *testing.T) {
	inputs := []TestData[string, []string]{
		{Title: "Should return an empty array on empty input", Input: "", Expected: []string{}},
		{Title: "Should return an array with three items", Input: filepath.Join("t1", "t2", "t3", "file.mp3"), Expected: []string{"t1", "t2", "t3"}},
		{Title: "Should return an array with one item", Input: filepath.Join("t1", "file.mp3"), Expected: []string{"t1"}},
		{Title: "Should return an empty array with only file name", Input: "file.mp3", Expected: []string{}},
	}

	for _, input := range inputs {
		t.Run(input.Title, func(t *testing.T) {
			file := FileTreeItem{
				Id:   "1",
				Path: input.Input,
			}
			result := file.GetPartsOfPath()

			if len(result) != len(input.Expected) {
				t.Errorf("The expected length '%v' differs from the received length '%v'.", len(input.Expected), len(result))
			}

			for i, _ := range result {
				if result[i] == input.Expected[i] {
					continue
				}
				t.Errorf("Expected '%v' in as value but received '%v'", input.Expected, result)
				break
			}
		})
	}
}

func Test_IsFileExtensionAllowed(t *testing.T) {
	file := FileTreeItem{
		Id:   "1",
		Path: filepath.Join("path", "file.mp3.vtt"),
	}
	allowedExtensions := []string{".mp4", ".mp3"}

	isAllowed := file.IsFileExtensionAllowed(allowedExtensions...)

	if isAllowed {
		t.Errorf("Expected false but received true")
	}
}

func Test_GetMimeType(t *testing.T) {
	testData := []TestData[string, string]{
		{Title: "Should return subtitle mime type", Input: "path/to/file.vtt", Expected: "text/vtt"},
		{Title: "Should return mp4 mime type", Input: "path/to/file.mp4", Expected: "video/mp4"},
		{Title: "Should return mpeg mime type", Input: "path/to/file.mp3", Expected: "audio/mpeg"},
		{Title: "Should return empty string", Input: "file", Expected: ""},
	}

	for _, data := range testData {

		t.Run(data.Title, func(t *testing.T) {
			// arrange
			file := FileTreeItem{
				Id:   "1",
				Path: data.Input,
			}

			// act
			mimeType := file.GetMimeType()

			// assert
			if mimeType != data.Expected {
				t.Errorf("Expected '%v' but received '%v'", data.Expected, mimeType)
			}
		})
	}
}
