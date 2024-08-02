package usecases

import (
	"backend/models"
	"fmt"
	"path/filepath"
	"testing"
)

func Test_IsFileExtensionAllowed_ValidMp4(t *testing.T) {
	fileItem := models.FileTreeItem{
		Id:   "1",
		Path: filepath.Join("path", "file.mp4"),
	}

	combinations := [][]string{
		{".mp4"},
		{".mp4", ".mp3"},
		{".mp4", ".vtt"},
		{".mp4", ".vtt", ".mp3"},
	}

	for _, combination := range combinations {
		t.Run(fmt.Sprintf("Should allow extensions: %v", combination), func(t *testing.T) {
			isAllowed := isFileExtensionAllowed(fileItem, combination...)

			if !isAllowed {
				t.Errorf("Expected true but received false")
			}
		})

	}
}

func Test_IsFileExtensionAllowed_ValidMp3(t *testing.T) {
	fileItem := models.FileTreeItem{
		Id:   "1",
		Path: filepath.Join("path", "file.mp3"),
	}

	combinations := [][]string{
		{".mp3"},
		{".mp4", ".mp3"},
		{".mp3", ".vtt"},
		{".mp4", ".vtt", ".mp3"},
	}

	for _, combination := range combinations {
		t.Run(fmt.Sprintf("Should allow extensions: %v", combination), func(t *testing.T) {
			isAllowed := isFileExtensionAllowed(fileItem, combination...)

			if !isAllowed {
				t.Errorf("Expected true but received false")
			}
		})

	}
}

func Test_IsFileExtensionAllowed_ValidVTT(t *testing.T) {
	fileItem := models.FileTreeItem{
		Id:   "1",
		Path: filepath.Join("path", "file.mp3.vtt"),
	}

	combinations := [][]string{
		{".vtt"},
		{".mp4", ".vtt"},
		{".mp3", ".vtt"},
		{".mp4", ".vtt", ".mp3"},
	}

	for _, combination := range combinations {
		t.Run(fmt.Sprintf("Should allow extensions: %v", combination), func(t *testing.T) {
			isAllowed := isFileExtensionAllowed(fileItem, combination...)

			if !isAllowed {
				t.Errorf("Expected true but received false")
			}
		})

	}
}

func Test_IsFileExtensionAllowed_Invalid(t *testing.T) {
	fileItem := models.FileTreeItem{
		Id:   "1",
		Path: filepath.Join("path", "file.mp3.vtt"),
	}
	allowedExtensions := []string{".mp4", ".mp3"}

	isAllowed := isFileExtensionAllowed(fileItem, allowedExtensions...)

	if isAllowed {
		t.Errorf("Expected false but received true")
	}
}
