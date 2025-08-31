package utilities

import (
	"backend/pkg/models"
	"testing"
)

func TestGetLowQualityImagePath(t *testing.T) {
	testCases := []models.TestData[string, string]{
		{
			Title:    "Simple image with .jpg extension",
			Input:    "image.jpg",
			Expected: "image_lowQuality.jpg",
		},
		{
			Title:    "Simple image with .png extension",
			Input:    "photo.png",
			Expected: "photo_lowQuality.png",
		},
		{
			Title:    "Image with .jpeg extension",
			Input:    "picture.jpeg",
			Expected: "picture_lowQuality.jpeg",
		},
		{
			Title:    "Image with .gif extension",
			Input:    "animation.gif",
			Expected: "animation_lowQuality.gif",
		},
		{
			Title:    "Image with .webp extension",
			Input:    "modern.webp",
			Expected: "modern_lowQuality.webp",
		},
		{
			Title:    "Image with .bmp extension",
			Input:    "bitmap.bmp",
			Expected: "bitmap_lowQuality.bmp",
		},
		{
			Title:    "Empty string",
			Input:    "",
			Expected: "",
		},
		{
			Title:    "File without extension",
			Input:    "noextension",
			Expected: "",
		},
		{
			Title:    "File with relative path",
			Input:    "images/photo.jpg",
			Expected: "images/photo_lowQuality.jpg",
		},
		{
			Title:    "File with deep nested path",
			Input:    "assets/images/gallery/photo.png",
			Expected: "assets/images/gallery/photo_lowQuality.png",
		},
		{
			Title:    "File with Unix-style absolute path",
			Input:    "/home/user/images/photo.jpg",
			Expected: "/home/user/images/photo_lowQuality.jpg",
		},
		{
			Title:    "File with Windows-style path",
			Input:    "C:\\Users\\user\\images\\photo.jpg",
			Expected: "C:\\Users\\user\\images\\photo_lowQuality.jpg",
		},
		{
			Title:    "Filename with spaces",
			Input:    "my photo.jpg",
			Expected: "my photo_lowQuality.jpg",
		},
		{
			Title:    "Filename with special characters",
			Input:    "image-test_01.jpg",
			Expected: "image-test_01_lowQuality.jpg",
		},
		{
			Title:    "Filename with Unicode characters",
			Input:    "测试图片.jpg",
			Expected: "测试图片_lowQuality.jpg",
		},
		{
			Title:    "Filename with umlauts",
			Input:    "bild_äöü.jpg",
			Expected: "bild_äöü_lowQuality.jpg",
		},
		{
			Title:    "Filename with multiple dots in name",
			Input:    "file.name.with.dots.jpg",
			Expected: "file.name.with.dots_lowQuality.jpg",
		},
		{
			Title:    "Hidden file starting with dot",
			Input:    ".hidden.jpg",
			Expected: ".hidden_lowQuality.jpg",
		},
		{
			Title:    "Just an extension",
			Input:    ".jpg",
			Expected: "_lowQuality.jpg",
		},
		{
			Title:    "Very long filename",
			Input:    "very_long_filename_that_might_cause_issues.jpg",
			Expected: "very_long_filename_that_might_cause_issues_lowQuality.jpg",
		},
		{
			Title:    "Filename with numbers",
			Input:    "image123.jpg",
			Expected: "image123_lowQuality.jpg",
		},
		{
			Title:    "Filename with uppercase extension",
			Input:    "image.JPG",
			Expected: "image_lowQuality.JPG",
		},
		{
			Title:    "Filename with mixed case extension",
			Input:    "PHOTO.Png",
			Expected: "PHOTO_lowQuality.Png",
		},
		{
			Title:    "Path with mixed separators",
			Input:    "path\\to/image.jpg",
			Expected: "path\\to/image_lowQuality.jpg",
		},
		{
			Title:    "Double extension file",
			Input:    "archive.tar.gz",
			Expected: "archive.tar_lowQuality.gz",
		},
		{
			Title:    "File with parentheses",
			Input:    "image(1).jpg",
			Expected: "image(1)_lowQuality.jpg",
		},
		{
			Title:    "File with brackets",
			Input:    "image[copy].jpg",
			Expected: "image[copy]_lowQuality.jpg",
		},
		{
			Title:    "Very short filename",
			Input:    "a.jpg",
			Expected: "a_lowQuality.jpg",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Title, func(t *testing.T) {
			result := GetLowQualityImagePath(testCase.Input)
			if result != testCase.Expected {
				t.Errorf("Expected %q, but got %q for input %q",
					testCase.Expected, result, testCase.Input)
			}
		})
	}
}
