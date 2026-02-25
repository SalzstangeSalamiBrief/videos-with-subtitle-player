package utilities

import (
	"backend/pkg/models"
	"testing"
)

type InputType struct {
	left  []models.FileTreeItem
	right []models.FileTreeItem
}

func TestGetDistinctFiles(t *testing.T) {
	testCases := []models.TestData[InputType, []models.FileTreeItem]{
		{
			Title: "Both slices empty - returns empty slice",
			Input: InputType{
				left:  []models.FileTreeItem{},
				right: []models.FileTreeItem{},
			},
			Expected: []models.FileTreeItem{},
		},
		{
			Title: "Left empty, right has items - returns empty slice",
			Input: InputType{
				left: []models.FileTreeItem{},
				right: []models.FileTreeItem{
					{Path: "/path/to/file1.txt"},
				},
			},
			Expected: []models.FileTreeItem{},
		},
		{
			Title: "Left has items, right empty - returns all left items",
			Input: InputType{
				left: []models.FileTreeItem{
					{Path: "/path/to/file1.txt"},
					{Path: "/path/to/file2.mp4"},
				},
				right: []models.FileTreeItem{},
			},
			Expected: []models.FileTreeItem{
				{Path: "/path/to/file1.txt"},
				{Path: "/path/to/file2.mp4"},
			},
		},
		{
			Title: "No overlap - returns all left items",
			Input: InputType{
				left: []models.FileTreeItem{
					{Path: "/path/to/file1.txt"},
					{Path: "/path/to/file2.mp4"},
				},
				right: []models.FileTreeItem{
					{Path: "/other/path/file3.json"},
					{Path: "/other/path/file4.go"},
				},
			},
			Expected: []models.FileTreeItem{
				{Path: "/path/to/file1.txt"},
				{Path: "/path/to/file2.mp4"},
			},
		},
		{
			Title: "Full overlap - returns empty slice",
			Input: InputType{
				left: []models.FileTreeItem{
					{Path: "/path/to/file1.txt"},
					{Path: "/path/to/file2.mp4"},
				},
				right: []models.FileTreeItem{
					{Path: "/path/to/file1.txt"},
					{Path: "/path/to/file2.mp4"},
				},
			},
			Expected: []models.FileTreeItem{},
		},
		{
			Title: "Partial overlap - returns only distinct items",
			Input: InputType{
				left: []models.FileTreeItem{
					{Path: "/path/to/file1.txt"},
					{Path: "/path/to/file2.mp4"},
					{Path: "/path/to/file3.json"},
				},
				right: []models.FileTreeItem{
					{Path: "/path/to/file2.mp4"},
				},
			},
			Expected: []models.FileTreeItem{
				{Path: "/path/to/file1.txt"},
				{Path: "/path/to/file3.json"},
			},
		},
		{
			Title: "Same name different path - returns item",
			Input: InputType{
				left: []models.FileTreeItem{
					{Path: "/path/a/file.txt"},
				},
				right: []models.FileTreeItem{
					{Path: "/path/b/file.txt"},
				},
			},
			Expected: []models.FileTreeItem{
				{Path: "/path/a/file.txt"},
			},
		},
		{
			Title: "Special characters in path - handles correctly",
			Input: InputType{
				left: []models.FileTreeItem{
					{Path: "/path/with spaces/file!@#$.txt"},
					{Path: "/path/日本語/ファイル.mp4"},
				},
				right: []models.FileTreeItem{
					{Path: "/path/with spaces/file!@#$.txt"},
				},
			},
			Expected: []models.FileTreeItem{
				{Path: "/path/日本語/ファイル.mp4"},
			},
		},
		{
			Title: "Single item in left, not in right - returns that item",
			Input: InputType{
				left: []models.FileTreeItem{
					{Path: "/unique/path/file.txt"},
				},
				right: []models.FileTreeItem{
					{Path: "/different/path/other.txt"},
				},
			},
			Expected: []models.FileTreeItem{
				{Path: "/unique/path/file.txt"},
			},
		},
		{
			Title: "Single item in left, exists in right - returns empty",
			Input: InputType{
				left: []models.FileTreeItem{
					{Path: "/path/file.txt"},
				},
				right: []models.FileTreeItem{
					{Path: "/path/file.txt"},
				},
			},
			Expected: []models.FileTreeItem{},
		},
		{
			Title: "Multiple dots in filename - handles correctly",
			Input: InputType{
				left: []models.FileTreeItem{
					{Path: "/path/my.multi.dot.file.tar.gz"},
				},
				right: []models.FileTreeItem{},
			},
			Expected: []models.FileTreeItem{
				{Path: "/path/my.multi.dot.file.tar.gz"},
			},
		},
		{
			Title: "Right has more items than left - returns distinct from left",
			Input: InputType{
				left: []models.FileTreeItem{
					{Path: "/path/file1.txt"},
				},
				right: []models.FileTreeItem{
					{Path: "/path/file1.txt"},
					{Path: "/path/file2.txt"},
					{Path: "/path/file3.txt"},
					{Path: "/path/file4.txt"},
				},
			},
			Expected: []models.FileTreeItem{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Title, func(t *testing.T) {
			result := GetDistinctFiles(testCase.Input.left, testCase.Input.right)

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
