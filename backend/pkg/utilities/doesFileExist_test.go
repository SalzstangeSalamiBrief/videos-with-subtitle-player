package utilities

import (
	"backend/pkg/models"
	"os"
	"path/filepath"
	"testing"
)

type InputType struct {
	subDir   string
	fileName string
}

var filePathsToCleanUp []string
var tempDir string

func beforeEach(t *testing.T) {
	tempDir = t.TempDir()
	testSubDir := filepath.Join(tempDir, "testSubDir")
	err := os.MkdirAll(testSubDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test dir %s: %v", testSubDir, err)
	}

	filePathsToCleanUp = []string{
		filepath.Join(tempDir, "simple.txt"),
		filepath.Join(tempDir, "video_file.mp4"),
		filepath.Join(tempDir, "my.multi.dot.file.json"),
		filepath.Join(tempDir, "file with spaces.docx"),
		filepath.Join(tempDir, "special-chars!@#$.log"),
		filepath.Join(tempDir, "UPPERCASE_FILE.PDF"),
		filepath.Join(tempDir, "mixed_Case-File.TXT"),
		filepath.Join(tempDir, "emoji_🎬_video.mkv"),
		filepath.Join(testSubDir, "nested_file.go"),
		filepath.Join(testSubDir, "another.test.spec.ts"),
		filepath.Join(testSubDir, "file_with_underscores_everywhere.md"),
		filepath.Join(testSubDir, "日本語ファイル.txt"),
		filepath.Join(testSubDir, "file(with)[brackets]{curly}.xml"),
		filepath.Join(testSubDir, ".hidden_file"),
		filepath.Join(testSubDir, "archive.tar.gz"),
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

func TestDoesFileExist(t *testing.T) {
	testCases := []models.TestData[InputType, bool]{
		{
			Title:    "File exists - simple.txt in root",
			Input:    InputType{subDir: "", fileName: "simple.txt"},
			Expected: true,
		},
		{
			Title:    "File exists - video_file.mp4 in root",
			Input:    InputType{subDir: "", fileName: "video_file.mp4"},
			Expected: true,
		},
		{
			Title:    "File exists - multiple dots in filename",
			Input:    InputType{subDir: "", fileName: "my.multi.dot.file.json"},
			Expected: true,
		},
		{
			Title:    "File exists - filename with spaces",
			Input:    InputType{subDir: "", fileName: "file with spaces.docx"},
			Expected: true,
		},
		{
			Title:    "File exists - special characters in filename",
			Input:    InputType{subDir: "", fileName: "special-chars!@#$.log"},
			Expected: true,
		},
		{
			Title:    "File exists - uppercase filename",
			Input:    InputType{subDir: "", fileName: "UPPERCASE_FILE.PDF"},
			Expected: true,
		},
		{
			Title:    "File exists - mixed case filename",
			Input:    InputType{subDir: "", fileName: "mixed_Case-File.TXT"},
			Expected: true,
		},
		{
			Title:    "File exists - emoji in filename",
			Input:    InputType{subDir: "", fileName: "emoji_🎬_video.mkv"},
			Expected: true,
		},
		{
			Title:    "File exists - nested_file.go in subdir",
			Input:    InputType{subDir: "testSubDir", fileName: "nested_file.go"},
			Expected: true,
		},
		{
			Title:    "File exists - multiple dots in subdir",
			Input:    InputType{subDir: "testSubDir", fileName: "another.test.spec.ts"},
			Expected: true,
		},
		{
			Title:    "File exists - underscores in filename in subdir",
			Input:    InputType{subDir: "testSubDir", fileName: "file_with_underscores_everywhere.md"},
			Expected: true,
		},
		{
			Title:    "File exists - Japanese characters in filename",
			Input:    InputType{subDir: "testSubDir", fileName: "日本語ファイル.txt"},
			Expected: true,
		},
		{
			Title:    "File exists - brackets in filename",
			Input:    InputType{subDir: "testSubDir", fileName: "file(with)[brackets]{curly}.xml"},
			Expected: true,
		},
		{
			Title:    "File exists - hidden file",
			Input:    InputType{subDir: "testSubDir", fileName: ".hidden_file"},
			Expected: true,
		},
		{
			Title:    "File exists - tar.gz extension",
			Input:    InputType{subDir: "testSubDir", fileName: "archive.tar.gz"},
			Expected: true,
		},
		{
			Title:    "File does not exist - nonexistent.txt",
			Input:    InputType{subDir: "", fileName: "nonexistent.txt"},
			Expected: false,
		},
		{
			Title:    "File does not exist - missing_file.mp4",
			Input:    InputType{subDir: "", fileName: "missing_file.mp4"},
			Expected: false,
		},
		{
			Title:    "File does not exist - does.not.exist.json in subdir",
			Input:    InputType{subDir: "testSubDir", fileName: "does.not.exist.json"},
			Expected: false,
		},
		{
			Title:    "File does not exist - special chars in subdir",
			Input:    InputType{subDir: "testSubDir", fileName: "fake_file_with_special!@#.log"},
			Expected: false,
		},
		{
			Title:    "File does not exist - another_missing_file.pdf",
			Input:    InputType{subDir: "", fileName: "another_missing_file.pdf"},
			Expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Title, func(t *testing.T) {
			beforeEach(t)

			pathUnderTest := filepath.Join(tempDir, testCase.Input.subDir, testCase.Input.fileName)
			result := DoesFileExist(pathUnderTest)
			if result != testCase.Expected {
				t.Errorf("Expected %v, but got %v for path: %s",
					testCase.Expected, result, pathUnderTest)
			}
		})
	}
}
