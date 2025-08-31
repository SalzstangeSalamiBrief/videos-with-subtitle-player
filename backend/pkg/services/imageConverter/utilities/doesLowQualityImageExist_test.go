package utilities

import (
	"backend/pkg/models"
	"os"
	"path/filepath"
	"testing"
)

type InputType struct {
	subDir          string
	sourceImagePath string
}

var filePathsToCleanUp []string
var tempDir string

func beforeEach(t *testing.T) {
	tempDir = t.TempDir()
	testSubDir := filepath.Join(tempDir, "subdir")
	err := os.MkdirAll(testSubDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test dir %s: %v", testSubDir, err)
	}

	filePathsToCleanUp = []string{
		filepath.Join(tempDir, "image_lowQuality.webp"),
		filepath.Join(testSubDir, "photo_lowQuality.webp"),
		filepath.Join(tempDir, "special-chars_äöü_lowQuality.webp"),
	}

	for _, file := range filePathsToCleanUp {
		f, err := os.Create(file)
		if err != nil {
			t.Fatalf("Failed to create test file %s: %v", file, err)
		}

		err = f.Close()
		if err != nil {
			t.Fatalf("Failed to close test file %s: %v", file, err)
		}
	}
}

func TestDoesLowQualityImageExist(t *testing.T) {
	testCases := []models.TestData[InputType, bool]{
		{
			Title: "File exists - simple case",
			Input: InputType{
				subDir:          "",
				sourceImagePath: "image.webp",
			},
			Expected: true,
		},
		{
			Title: "File does not exist",
			Input: InputType{
				subDir:          "",
				sourceImagePath: "nonexistent.jpg",
			},
			Expected: false,
		},
		{
			Title: "File does not exist - different extension",
			Input: InputType{
				subDir:          "",
				sourceImagePath: "image.gif",
			},
			Expected: false,
		},
		{
			Title: "Empty source image path",
			Input: InputType{
				subDir:          "",
				sourceImagePath: "",
			},
			Expected: false,
		},
		{
			Title: "Source path without extension",
			Input: InputType{
				subDir:          "",
				sourceImagePath: "noextension",
			},
			Expected: false,
		},
		{
			Title: "Root path does not exist",
			Input: InputType{
				subDir:          filepath.Join("", "nonexistent"),
				sourceImagePath: "image.jpg",
			},
			Expected: false,
		},
		{
			Title: "Empty root path",
			Input: InputType{
				subDir:          "",
				sourceImagePath: "image.jpg",
			},
			Expected: false,
		},
		{
			Title: "Both paths empty",
			Input: InputType{
				subDir:          "",
				sourceImagePath: "",
			},
			Expected: false,
		},
		{
			Title: "Path with multiple dots",
			Input: InputType{
				subDir:          "",
				sourceImagePath: "file.name.with.dots.jpg",
			},
			Expected: false,
		},
		{
			Title: "Path traversal attempt - relative path",
			Input: InputType{
				subDir:          "",
				sourceImagePath: "../image.jpg",
			},
			Expected: false,
		},
		{
			Title: "Very long filename without counter part",
			Input: InputType{
				subDir:          "",
				sourceImagePath: "very_long_filename_that_might_cause_issues_in_some_filesystems_or_scenarios_but_should_still_work_correctly.jpg",
			},
			Expected: false,
		},
		{
			Title: "File without extension",
			Input: InputType{
				subDir:          "",
				sourceImagePath: "image",
			},
			Expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Title, func(t *testing.T) {
			beforeEach(t)

			pathUnderTest := filepath.Join(tempDir, testCase.Input.subDir)
			result := DoesLowQualityImageExist(pathUnderTest, testCase.Input.sourceImagePath)
			if result != testCase.Expected {
				t.Errorf("Expected %v, but got %v for subDir: %s, sourceImagePath: %s",
					testCase.Expected, result, pathUnderTest, testCase.Input.sourceImagePath)
			}
		})
	}
}
