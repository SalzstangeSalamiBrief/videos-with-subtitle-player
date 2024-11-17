package models

import "testing"

type ReceiverInput struct {
	Entity FileTreeDto
	Param  string
}

func Test_GetIndexOfChildByName(t *testing.T) {
	var inputs = []TestData[ReceiverInput, int]{
		{
			Title: "Should return -1 for empty children",
			Input: ReceiverInput{
				Entity: FileTreeDto{
					Name:        "",
					Id:          "",
					ThumbnailId: "",
					Files:       nil,
					Children:    nil,
				},
				Param: "",
			},
			Expected: -1,
		},
		{
			Title: "Should return -1 for empty input",
			Input: ReceiverInput{
				Entity: FileTreeDto{
					Name:        "",
					Id:          "",
					ThumbnailId: "",
					Files:       nil,
					Children: []FileTreeDto{
						{
							Name:        "a",
							Id:          "a",
							ThumbnailId: "a",
							Files:       nil,
							Children:    nil,
						},
					},
				},
				Param: "",
			},
			Expected: -1,
		},
		{
			Title: "Should return -1 for invalid input",
			Input: ReceiverInput{
				Entity: FileTreeDto{
					Name:        "",
					Id:          "",
					ThumbnailId: "",
					Files:       nil,
					Children: []FileTreeDto{
						{
							Name:        "a",
							Id:          "a",
							ThumbnailId: "a",
							Files:       nil,
							Children:    nil,
						},
					},
				},
				Param: "b",
			},
			Expected: -1,
		},
		{
			Title: "Should return 0 on match for 1 child",
			Input: ReceiverInput{
				Entity: FileTreeDto{
					Name:        "",
					Id:          "",
					ThumbnailId: "",
					Files:       nil,
					Children: []FileTreeDto{
						{
							Name:        "a",
							Id:          "a",
							ThumbnailId: "a",
							Files:       nil,
							Children:    nil,
						},
					},
				},
				Param: "a",
			},
			Expected: 0,
		},
		{
			Title: "Should return 0 on match for 2 children",
			Input: ReceiverInput{
				Entity: FileTreeDto{
					Name:        "",
					Id:          "",
					ThumbnailId: "",
					Files:       nil,
					Children: []FileTreeDto{
						{
							Name:        "a",
							Id:          "a",
							ThumbnailId: "a",
							Files:       nil,
							Children:    nil,
						},
						{
							Name:        "b",
							Id:          "b",
							ThumbnailId: "b",
							Files:       nil,
							Children:    nil,
						},
					},
				},
				Param: "a",
			},
			Expected: 0,
		},
		{
			Title: "Should return 1 on match for 2 children",
			Input: ReceiverInput{
				Entity: FileTreeDto{
					Name:        "",
					Id:          "",
					ThumbnailId: "",
					Files:       nil,
					Children: []FileTreeDto{
						{
							Name:        "a",
							Id:          "a",
							ThumbnailId: "a",
							Files:       nil,
							Children:    nil,
						},
						{
							Name:        "b",
							Id:          "b",
							ThumbnailId: "b",
							Files:       nil,
							Children:    nil,
						},
					},
				},
				Param: "b",
			},
			Expected: 1,
		},
	}

	for _, input := range inputs {
		t.Run(input.Title, func(t *testing.T) {
			result := input.Input.Entity.GetIndexOfChildByName(input.Input.Param)
			if result != input.Expected {
				t.Errorf("The expected result '%v' differs from the received value '%v'.", input.Expected, result)
			}
		})
	}
}
