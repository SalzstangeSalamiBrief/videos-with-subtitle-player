package models

import (
	"backend/pkg/models"
	"reflect"
	"testing"
)

func TestConvertImageMagickCommandsArrayToArgumentsArray(t *testing.T) {
	inputs := []models.TestData[[]ImageCLICommand, []string]{
		{
			Title:    "Empty array returns empty slice",
			Input:    []ImageCLICommand{},
			Expected: []string{},
		},
		{
			Title: "Single command with both Command and Arg",
			Input: []ImageCLICommand{
				{Command: "-resize", Arg: "100x100"},
			},
			Expected: []string{"-resize", "100x100"},
		},
		{
			Title: "Single command with only Command",
			Input: []ImageCLICommand{
				{Command: "-auto-orient", Arg: ""},
			},
			Expected: []string{"-auto-orient"},
		},
		{
			Title: "Single command with only Arg",
			Input: []ImageCLICommand{
				{Command: "", Arg: "output.jpg"},
			},
			Expected: []string{"output.jpg"},
		},
		{
			Title: "Single command with empty Command and Arg",
			Input: []ImageCLICommand{
				{Command: "", Arg: ""},
			},
			Expected: []string{},
		},
		{
			Title: "Multiple commands with both Command and Arg",
			Input: []ImageCLICommand{
				{Command: "-resize", Arg: "100x100"},
				{Command: "-quality", Arg: "90"},
				{Command: "-format", Arg: "webp"},
			},
			Expected: []string{"-resize", "100x100", "-quality", "90", "-format", "webp"},
		},
		{
			Title: "Multiple commands with mixed patterns",
			Input: []ImageCLICommand{
				{Command: "-resize", Arg: "100x100"},
				{Command: "-auto-orient", Arg: ""},
				{Command: "", Arg: "input.jpg"},
				{Command: "-quality", Arg: "85"},
				{Command: "", Arg: ""},
				{Command: "", Arg: "output.webp"},
			},
			Expected: []string{"-resize", "100x100", "-auto-orient", "input.jpg", "-quality", "85", "output.webp"},
		},
		{
			Title: "Commands with whitespace values",
			Input: []ImageCLICommand{
				{Command: "-resize", Arg: "100x100"},
				{Command: "-comment", Arg: "This is a test image"},
			},
			Expected: []string{"-resize", "100x100", "-comment", "This is a test image"},
		},
		{
			Title: "Commands with special characters",
			Input: []ImageCLICommand{
				{Command: "-resize", Arg: "50%"},
				{Command: "-gravity", Arg: "center"},
				{Command: "-extent", Arg: "200x200+10+10"},
			},
			Expected: []string{"-resize", "50%", "-gravity", "center", "-extent", "200x200+10+10"},
		},
		{
			Title: "Real-world WebP conversion example",
			Input: []ImageCLICommand{
				{Command: "", Arg: "input.jpg"},
				{Command: "-quality", Arg: "80"},
				{Command: "-define", Arg: "webp:lossless=false"},
				{Command: "-resize", Arg: "1920x1080>"},
				{Command: "-auto-orient", Arg: ""},
				{Command: "", Arg: "output.webp"},
			},
			Expected: []string{"input.jpg", "-quality", "80", "-define", "webp:lossless=false", "-resize", "1920x1080>", "-auto-orient", "output.webp"},
		},
	}

	for _, input := range inputs {
		t.Run(input.Title, func(t *testing.T) {
			result := ConvertImageMagickCommandsArrayToArgumentsArray(input.Input)

			if !reflect.DeepEqual(result, input.Expected) {
				t.Errorf("Expected %v but received %v", input.Expected, result)
			}
		})
	}
}
