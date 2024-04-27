package usecases

import (
	"backend/models"
	"fmt"
	"testing"
)

func Test_getPartsOfPath(t *testing.T) {
	inputs := []models.TestData[string, []string]{
		{Input: "", Expected: []string{}},
		{Input: "t1\\t2\\t3\\file.mp3", Expected: []string{"t1", "t2", "t3"}},
		{Input: "t1\\file.mp3", Expected: []string{"t1"}},
		{Input: "file.mp3", Expected: []string{}},
	}

	for _, input := range inputs {
		t.Run(fmt.Sprintf("input: %v", input.Input), func(t *testing.T) {
			file := models.FileTreeItem{
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
