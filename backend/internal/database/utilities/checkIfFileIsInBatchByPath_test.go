package utilities

import (
	"backend/pkg/models"
	"testing"
)

type checkIfFileIsInBatchInput struct {
	path  string
	batch []models.FileTreeItem
}

func TestCheckIfFileIsInBatchByPath(t *testing.T) {
	testCases := []models.TestData[checkIfFileIsInBatchInput, bool]{
		{
			Title: "Returns true when file exists in batch with exact path match",
			Input: checkIfFileIsInBatchInput{
				path: "/videos/movie.mp4",
				batch: []models.FileTreeItem{
					{Path: "/videos/movie.mp4"},
					{Path: "/videos/trailer.mp4"},
				},
			},
			Expected: true,
		},
		{
			Title: "Returns false when file does not exist in batch",
			Input: checkIfFileIsInBatchInput{
				path: "/videos/missing.mp4",
				batch: []models.FileTreeItem{
					{Path: "/videos/movie.mp4"},
					{Path: "/videos/trailer.mp4"},
				},
			},
			Expected: false,
		},
		{
			Title: "Returns false when batch is empty",
			Input: checkIfFileIsInBatchInput{
				path:  "/videos/movie.mp4",
				batch: []models.FileTreeItem{},
			},
			Expected: false,
		},
		{
			Title: "Returns true when file is the only item in batch",
			Input: checkIfFileIsInBatchInput{
				path: "/single/file.txt",
				batch: []models.FileTreeItem{
					{Path: "/single/file.txt"},
				},
			},
			Expected: true,
		},
		{
			Title: "Returns false when path partially matches but is not exact",
			Input: checkIfFileIsInBatchInput{
				path: "/videos/movie",
				batch: []models.FileTreeItem{
					{Path: "/videos/movie.mp4"},
				},
			},
			Expected: false,
		},
		{
			Title: "Returns false when path is a relative path without a file",
			Input: checkIfFileIsInBatchInput{
				path: "/videos",
				batch: []models.FileTreeItem{
					{Path: "/videos/movie.mp4"},
				},
			},
			Expected: false,
		},
		{
			Title: "Returns true when path matches first item in batch",
			Input: checkIfFileIsInBatchInput{
				path: "/first/item.mp4",
				batch: []models.FileTreeItem{
					{Path: "/first/item.mp4"},
					{Path: "/second/item.mp4"},
					{Path: "/third/item.mp4"},
				},
			},
			Expected: true,
		},
		{
			Title: "Returns true when path matches last item in batch",
			Input: checkIfFileIsInBatchInput{
				path: "/third/item.mp4",
				batch: []models.FileTreeItem{
					{Path: "/first/item.mp4"},
					{Path: "/second/item.mp4"},
					{Path: "/third/item.mp4"},
				},
			},
			Expected: true,
		},
		{
			Title: "Returns false for empty path",
			Input: checkIfFileIsInBatchInput{
				path: "",
				batch: []models.FileTreeItem{
					{Path: "/videos/movie.mp4"},
				},
			},
			Expected: false,
		},
		{
			Title: "Returns true when path contains special characters",
			Input: checkIfFileIsInBatchInput{
				path: "/videos/my file (2024) [HD].mp4",
				batch: []models.FileTreeItem{
					{Path: "/videos/my file (2024) [HD].mp4", Name: "my file (2024) [HD].mp4"},
				},
			},
			Expected: true,
		},
		{
			Title: "Returns false when path case differs",
			Input: checkIfFileIsInBatchInput{
				path: "/Videos/Movie.MP4",
				batch: []models.FileTreeItem{
					{Path: "/videos/movie.mp4"},
				},
			},
			Expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Title, func(t *testing.T) {
			result := CheckIfFileIsInBatchByPath(testCase.Input.path, testCase.Input.batch)
			if result != testCase.Expected {
				t.Errorf("Expected %v, got %v", testCase.Expected, result)
			}
		})
	}
}
