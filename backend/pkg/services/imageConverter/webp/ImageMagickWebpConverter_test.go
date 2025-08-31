package webp

import (
	"backend/pkg/models"
	"path/filepath"
	"testing"
)

type ContainsLowQualityImagePathTestInput struct {
	RelativeImagePath string
	AllImagePaths     []string
}

func TestContainsLowQualityImagePath(t *testing.T) {
	testCases := []models.TestData[ContainsLowQualityImagePathTestInput, bool]{
		{
			Title: "Low quality image exists in list",
			Input: ContainsLowQualityImagePathTestInput{
				RelativeImagePath: "image.jpg",
				AllImagePaths:     []string{"image.jpg", "image_lowQuality.jpg", "other.png"},
			},
			Expected: true,
		},
		{
			Title: "Low quality image does not exist in list",
			Input: ContainsLowQualityImagePathTestInput{
				RelativeImagePath: "image.jpg",
				AllImagePaths:     []string{"image.jpg", "other.png", "another.webp"},
			},
			Expected: false,
		},
		{
			Title: "Empty image paths list",
			Input: ContainsLowQualityImagePathTestInput{
				RelativeImagePath: "image.jpg",
				AllImagePaths:     []string{},
			},
			Expected: false,
		},
		{
			Title: "Image with path - low quality exists",
			Input: ContainsLowQualityImagePathTestInput{
				RelativeImagePath: filepath.Join("folder", "subfolder", "photo.png"),
				AllImagePaths:     []string{filepath.Join("folder", "subfolder", "photo.png"), filepath.Join("folder", "subfolder", "photo_lowQuality.png")},
			},
			Expected: true,
		},
		{
			Title: "WebP image - low quality exists",
			Input: ContainsLowQualityImagePathTestInput{
				RelativeImagePath: "image.webp",
				AllImagePaths:     []string{"image.webp", "image_lowQuality.webp", "test.jpg"},
			},
			Expected: true,
		},
		{
			Title: "Multiple images but no matching low quality",
			Input: ContainsLowQualityImagePathTestInput{
				RelativeImagePath: "target.jpg",
				AllImagePaths:     []string{"other_lowQuality.jpg", "different_lowQuality.png", "target.jpg"},
			},
			Expected: false,
		},
		{
			Title: "Exact match with low quality naming pattern",
			Input: ContainsLowQualityImagePathTestInput{
				RelativeImagePath: filepath.Join("assets", "gallery", "photo.jpeg"),
				AllImagePaths:     []string{filepath.Join("assets", "gallery", "photo.jpeg"), filepath.Join("assets", "gallery", "helloWorld.jpeg"), filepath.Join("assets", "gallery", "helloWorld_lowQuality.jpeg"), filepath.Join("assets", "gallery", "dummy.jpeg"), filepath.Join("assets", "gallery", "photo_lowQuality.jpeg")},
			},
			Expected: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Title, func(t *testing.T) {
			result := containsLowQualityImagePath(testCase.Input.RelativeImagePath, testCase.Input.AllImagePaths)
			if result != testCase.Expected {
				t.Errorf("Expected %v, but got %v for input %+v",
					testCase.Expected, result, testCase.Input)
			}
		})
	}
}
