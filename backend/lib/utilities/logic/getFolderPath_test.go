package logic

import (
	"backend/models"
	"path/filepath"
	"testing"
)

func Test_GetFolderPath(t *testing.T) {
	rootPath := filepath.Join("C:", "test")
	inputs := []models.TestData[GetFolderPathInput, string]{
		{Title: "Should return an empty string on empty inputs", Input: GetFolderPathInput{Path: "", RootPath: ""}, Expected: ""},
		{Title: "Should return an empty string on empty path input", Input: GetFolderPathInput{Path: "", RootPath: rootPath}, Expected: ""},
		{Title: "Should return a path based on input without file", Input: GetFolderPathInput{Path: filepath.Join("D:", "hello_world", "test"), RootPath: rootPath}, Expected: filepath.Join("D:", "hello_world", "test")},
		{Title: "Should return a path based on input with file", Input: GetFolderPathInput{Path: filepath.Join("D:", "hello_world", "test", "file.md"), RootPath: rootPath}, Expected: filepath.Join("D:", "hello_world", "test", "file.md")},
		{Title: "Should return an empty string on equal path and root path", Input: GetFolderPathInput{Path: rootPath, RootPath: rootPath}, Expected: ""},
		{Title: "Should return path with one segment", Input: GetFolderPathInput{Path: filepath.Join(rootPath, "hello_world"), RootPath: rootPath}, Expected: filepath.Join(string(filepath.Separator), "hello_world")},
		{Title: "Should return path with one segment and file name", Input: GetFolderPathInput{Path: filepath.Join(rootPath, "hello_world", "file.md"), RootPath: rootPath}, Expected: filepath.Join(string(filepath.Separator), "hello_world", "file.md")},
	}

	for _, input := range inputs {
		t.Run(input.Title, func(t *testing.T) {
			result := GetFolderPath(input.Input)

			if result != input.Expected {
				t.Errorf("Expected '%v' but received '%v'", input.Expected, result)
			}
		})
	}
}
