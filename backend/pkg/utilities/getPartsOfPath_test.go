package utilities

import (
	models2 "backend/pkg/models"
	"path/filepath"
	"testing"
)

func Test_GetPartsOfPath(t *testing.T) {
	inputs := []models2.TestData[string, []string]{
		{Title: "Should return an empty array on empty input", Input: "", Expected: []string{}},
		{Title: "Should return an array with three items", Input: filepath.Join("t1", "t2", "t3", "file.mp3"), Expected: []string{"t1", "t2", "t3"}},
		{Title: "Should return an array with one item", Input: filepath.Join("t1", "file.mp3"), Expected: []string{"t1"}},
		{Title: "Should return an empty array with only file name", Input: "file.mp3", Expected: []string{}},
	}

	for _, input := range inputs {
		t.Run(input.Title, func(t *testing.T) {
			file := models2.FileTreeItem{
				Id:   "1",
				Path: input.Input,
			}
			result := GetPartsOfPath(file)

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
