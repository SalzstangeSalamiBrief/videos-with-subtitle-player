package database

import (
	"backend/pkg/models"
	"testing"
)

type InputType struct {
	left  []models.FileNode
	right []models.FileNode
}

func TestGetDistinctFiles(t *testing.T) {
	testCases := []models.TestData[InputType, []models.FileNode]{
		{
			Title: "Both slices empty - returns empty slice",
			Input: InputType{
				left:  []models.FileNode{},
				right: []models.FileNode{},
			},
			Expected: []models.FileNode{},
		},
		{
			Title: "Left empty, right has items - returns empty slice",
			Input: InputType{
				left: []models.FileNode{},
				right: []models.FileNode{
					{Path: "/path/to/file1.txt"},
				},
			},
			Expected: []models.FileNode{},
		},
		{
			Title: "Left has items, right empty - returns all left items",
			Input: InputType{
				left: []models.FileNode{
					{Path: "/path/to/file1.txt"},
					{Path: "/path/to/file2.mp4"},
				},
				right: []models.FileNode{},
			},
			Expected: []models.FileNode{
				{Path: "/path/to/file1.txt"},
				{Path: "/path/to/file2.mp4"},
			},
		},
		{
			Title: "No overlap - returns all left items",
			Input: InputType{
				left: []models.FileNode{
					{Path: "/path/to/file1.txt"},
					{Path: "/path/to/file2.mp4"},
				},
				right: []models.FileNode{
					{Path: "/other/path/file3.json"},
					{Path: "/other/path/file4.go"},
				},
			},
			Expected: []models.FileNode{
				{Path: "/path/to/file1.txt"},
				{Path: "/path/to/file2.mp4"},
			},
		},
		{
			Title: "Full overlap - returns empty slice",
			Input: InputType{
				left: []models.FileNode{
					{Path: "/path/to/file1.txt"},
					{Path: "/path/to/file2.mp4"},
				},
				right: []models.FileNode{
					{Path: "/path/to/file1.txt"},
					{Path: "/path/to/file2.mp4"},
				},
			},
			Expected: []models.FileNode{},
		},
		{
			Title: "Partial overlap - returns only distinct items",
			Input: InputType{
				left: []models.FileNode{
					{Path: "/path/to/file1.txt"},
					{Path: "/path/to/file2.mp4"},
					{Path: "/path/to/file3.json"},
				},
				right: []models.FileNode{
					{Path: "/path/to/file2.mp4"},
				},
			},
			Expected: []models.FileNode{
				{Path: "/path/to/file1.txt"},
				{Path: "/path/to/file3.json"},
			},
		},
		{
			Title: "Same name different path - returns item",
			Input: InputType{
				left: []models.FileNode{
					{Path: "/path/a/file.txt"},
				},
				right: []models.FileNode{
					{Path: "/path/b/file.txt"},
				},
			},
			Expected: []models.FileNode{
				{Path: "/path/a/file.txt"},
			},
		},
		{
			Title: "Special characters in path - handles correctly",
			Input: InputType{
				left: []models.FileNode{
					{Path: "/path/with spaces/file!@#$.txt"},
					{Path: "/path/日本語/ファイル.mp4"},
				},
				right: []models.FileNode{
					{Path: "/path/with spaces/file!@#$.txt"},
				},
			},
			Expected: []models.FileNode{
				{Path: "/path/日本語/ファイル.mp4"},
			},
		},
		{
			Title: "Single item in left, not in right - returns that item",
			Input: InputType{
				left: []models.FileNode{
					{Path: "/unique/path/file.txt"},
				},
				right: []models.FileNode{
					{Path: "/different/path/other.txt"},
				},
			},
			Expected: []models.FileNode{
				{Path: "/unique/path/file.txt"},
			},
		},
		{
			Title: "Single item in left, exists in right - returns empty",
			Input: InputType{
				left: []models.FileNode{
					{Path: "/path/file.txt"},
				},
				right: []models.FileNode{
					{Path: "/path/file.txt"},
				},
			},
			Expected: []models.FileNode{},
		},
		{
			Title: "Multiple dots in filename - handles correctly",
			Input: InputType{
				left: []models.FileNode{
					{Path: "/path/my.multi.dot.file.tar.gz"},
				},
				right: []models.FileNode{},
			},
			Expected: []models.FileNode{
				{Path: "/path/my.multi.dot.file.tar.gz"},
			},
		},
		{
			Title: "Right has more items than left - returns distinct from left",
			Input: InputType{
				left: []models.FileNode{
					{Path: "/path/file1.txt"},
				},
				right: []models.FileNode{
					{Path: "/path/file1.txt"},
					{Path: "/path/file2.txt"},
					{Path: "/path/file3.txt"},
					{Path: "/path/file4.txt"},
				},
			},
			Expected: []models.FileNode{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Title, func(t *testing.T) {
			result := getDistinctFiles(testCase.Input.left, testCase.Input.right)

			if len(result) != len(testCase.Expected) {
				t.Errorf("Expected %d items, but got %d items",
					len(testCase.Expected), len(result))
				return
			}

			for i, expectedItem := range testCase.Expected {
				if result[i].Path != expectedItem.Path {
					t.Errorf("Expected path %s at index %d, but got %s",
						expectedItem.Path, i, result[i].Path)
				}
			}
		})
	}
}
