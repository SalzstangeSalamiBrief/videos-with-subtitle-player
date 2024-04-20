package lib

import (
	"backend/models"
	"fmt"
	"path/filepath"
	"testing"
)

func Test_GetFolderPath(t *testing.T) {
	rootPath := filepath.Join("C:", "test")
	inputs := []models.TestData[GetFolderPathInput, string]{
		{Input: GetFolderPathInput{Path: "", RootPath: ""}, Expected: ""},
		{Input: GetFolderPathInput{Path: "", RootPath: rootPath}, Expected: ""},
		{Input: GetFolderPathInput{Path: filepath.Join("D:", "hello_world", "test"), RootPath: rootPath}, Expected: filepath.Join("D:", "hello_world", "test")},
		{Input: GetFolderPathInput{Path: filepath.Join("D:", "hello_world", "test", "file.md"), RootPath: rootPath}, Expected: filepath.Join("D:", "hello_world", "test", "file.md")},
		{Input: GetFolderPathInput{Path: filepath.Join("D:", "hello_world", "test", "file.md"), RootPath: rootPath}, Expected: filepath.Join("D:", "hello_world", "test", "file.md")},
		{Input: GetFolderPathInput{Path: rootPath, RootPath: rootPath}, Expected: ""},
		{Input: GetFolderPathInput{Path: filepath.Join(rootPath, "hello_world"), RootPath: rootPath}, Expected: filepath.Join("\\hello_world")},
		{Input: GetFolderPathInput{Path: filepath.Join(rootPath, "hello_world", "file.md"), RootPath: rootPath}, Expected: filepath.Join("\\hello_world", "file.md")},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("input: '%v'", input.Input), func(t *testing.T) {
			result := GetFolderPath(input.Input)

			if result != input.Expected {
				t.Errorf("Expected '%v' but received '%v'", input.Expected, result)
			}
		})
	}
}
