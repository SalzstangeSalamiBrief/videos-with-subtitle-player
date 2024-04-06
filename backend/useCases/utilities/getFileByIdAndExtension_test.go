package utilities

import (
	"backend/models"
	"fmt"
	"testing"
)

func TestIsFileExtensionAllowedValidMp4(t *testing.T) {
	fileItem := models.FileTreeItem{
		Id:   "1",
		Path: "path/to/file.mp4",
	}

	combinations := [][]string{
		{".mp4"},
		{".mp4", ".mp3"},
		{".mp4", ".vtt"},
		{".mp4", ".vtt", ".mp3"},
	}

	for _, combination := range combinations {
		t.Run(fmt.Sprintf("allowed extensions: %v", combination), func(t *testing.T) {
			// act
			isAllowed := isFileExtensionAllowed(fileItem, combination...)

			// assert
			if !isAllowed {
				t.Errorf("Expected true but received false")
			}
		})

	}
}

func TestIsFileExtensionAllowedValidMp3(t *testing.T) {
	fileItem := models.FileTreeItem{
		Id:   "1",
		Path: "path/to/file.mp3",
	}

	combinations := [][]string{
		{".mp3"},
		{".mp4", ".mp3"},
		{".mp3", ".vtt"},
		{".mp4", ".vtt", ".mp3"},
	}

	for _, combination := range combinations {
		t.Run(fmt.Sprintf("allowed extensions: %v", combination), func(t *testing.T) {
			// act
			isAllowed := isFileExtensionAllowed(fileItem, combination...)

			// assert
			if !isAllowed {
				t.Errorf("Expected true but received false")
			}
		})

	}
}

func TestIsFileExtensionAllowedValidVTT(t *testing.T) {
	fileItem := models.FileTreeItem{
		Id:   "1",
		Path: "path/to/file.mp3.vtt",
	}

	combinations := [][]string{
		{".vtt"},
		{".mp4", ".vtt"},
		{".mp3", ".vtt"},
		{".mp4", ".vtt", ".mp3"},
	}

	for _, combination := range combinations {
		t.Run(fmt.Sprintf("allowed extensions: %v", combination), func(t *testing.T) {
			// act
			isAllowed := isFileExtensionAllowed(fileItem, combination...)

			// assert
			if !isAllowed {
				t.Errorf("Expected true but received false")
			}
		})

	}
}

func TestIsFileExtensionAllowedInvalid(t *testing.T) {
	fileItem := models.FileTreeItem{
		Id:   "1",
		Path: "path/to/file.mp3.vtt",
	}
	allowedExtensions := []string{".mp4", ".mp3"}

	// act
	isAllowed := isFileExtensionAllowed(fileItem, allowedExtensions...)

	// assert
	if isAllowed {
		t.Errorf("Expected false but received true")
	}
}
