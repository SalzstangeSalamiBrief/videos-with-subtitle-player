package utilities

import (
	"backend/pkg/models"
	"testing"
)

func TestIsWebpImage(t *testing.T) {
	testCases := []models.TestData[string, bool]{
		{
			Title:    "WebP image with .webp extension",
			Input:    "image.webp",
			Expected: true,
		},
		{
			Title:    "WebP image with path",
			Input:    "folder/image.webp",
			Expected: true,
		},
		{
			Title:    "WebP image with deep path",
			Input:    "assets/images/gallery/photo.webp",
			Expected: true,
		},
		{
			Title:    "WebP image with absolute path",
			Input:    "/home/user/images/picture.webp",
			Expected: true,
		},
		{
			Title:    "WebP image with Windows path",
			Input:    "C:\\Users\\user\\images\\photo.webp",
			Expected: true,
		},
		{
			Title:    "JPEG image",
			Input:    "image.jpg",
			Expected: false,
		},
		{
			Title:    "JPEG image with .jpeg extension",
			Input:    "photo.jpeg",
			Expected: false,
		},
		{
			Title:    "PNG image",
			Input:    "image.png",
			Expected: false,
		},
		{
			Title:    "GIF image",
			Input:    "animation.gif",
			Expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Title, func(t *testing.T) {
			result := IsWebpImage(testCase.Input)
			if result != testCase.Expected {
				t.Errorf("Expected %v, but got %v for input %q",
					testCase.Expected, result, testCase.Input)
			}
		})
	}
}
