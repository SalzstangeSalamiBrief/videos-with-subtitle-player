package utilities

import (
	"backend/pkg/models"
	"testing"
)

func TestIsLossyImage(t *testing.T) {
	testCases := []models.TestData[string, bool]{
		{
			Title:    "JPEG image with .jpg extension",
			Input:    "image.jpg",
			Expected: true,
		},
		{
			Title:    "JPEG image with .jpeg extension",
			Input:    "photo.jpeg",
			Expected: true,
		},
		{
			Title:    "JPEG image with path",
			Input:    "folder/image.jpg",
			Expected: true,
		},
		{
			Title:    "JPEG image with deep path",
			Input:    "assets/images/gallery/photo.jpeg",
			Expected: true,
		},
		{
			Title:    "JPEG image with absolute path",
			Input:    "/home/user/images/picture.jpg",
			Expected: true,
		},
		{
			Title:    "JPEG image with Windows path",
			Input:    "C:\\Users\\user\\images\\photo.jpeg",
			Expected: true,
		},
		{
			Title:    "PNG image - not lossy",
			Input:    "image.png",
			Expected: false,
		},
		{
			Title:    "WebP image - not lossy (can be lossless)",
			Input:    "image.webp",
			Expected: false,
		},
		{
			Title:    "GIF image - not lossy",
			Input:    "animation.gif",
			Expected: false,
		},

		{
			Title:    "Empty string",
			Input:    "",
			Expected: false,
		},
		{
			Title:    "Just the extension .jpg",
			Input:    ".jpg",
			Expected: true,
		},
		{
			Title:    "Just the extension .jpeg",
			Input:    ".jpeg",
			Expected: true,
		},

		{
			Title:    "Filename with spaces",
			Input:    "my photo.jpg",
			Expected: true,
		},
		{
			Title:    "Filename with special characters",
			Input:    "image-test_01.jpeg",
			Expected: true,
		},
		{
			Title:    "Filename with Unicode characters",
			Input:    "测试图片.jpg",
			Expected: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Title, func(t *testing.T) {
			result := IsLossyImage(testCase.Input)
			if result != testCase.Expected {
				t.Errorf("Expected %v, but got %v for input %q",
					testCase.Expected, result, testCase.Input)
			}
		})
	}
}
