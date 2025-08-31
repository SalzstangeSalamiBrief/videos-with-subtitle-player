package utilities

import (
	"backend/pkg/models"
	"testing"
)

func TestGetLowQualityImageName(t *testing.T) {
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
			Title:    "Image with .tiff extension",
			Input:    "professional.tiff",
			Expected: "professional_lowQuality.tiff",
		},
		{
			Title:    "Empty string",
			Input:    "",
			Expected: "",
		},
		{
			Title:    "File with path - Unix style",
			Input:    "/home/user/images/photo.jpg",
			Expected: "photo_lowQuality.jpg",
		},
		{
			Title:    "File with path - Windows style",
			Input:    "C:\\Users\\user\\images\\photo.jpg",
			Expected: "photo_lowQuality.jpg",
		},
		{
			Title:    "File with relative path",
			Input:    "images/subfolder/picture.png",
			Expected: "picture_lowQuality.png",
		},
		{
			Title:    "File with deep nested path",
			Input:    "very/deep/nested/folder/structure/image.jpg",
			Expected: "image_lowQuality.jpg",
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
			Title:    "Filename with multiple dots",
			Input:    "file.name.with.dots.jpg",
			Expected: "file.name.with.dots_lowQuality.jpg",
		},
		{
			Title:    "Filename starting with dot (hidden file)",
			Input:    ".hidden.jpg",
			Expected: ".hidden_lowQuality.jpg",
		},
		{
			Title:    "Filename with only dot and extension",
			Input:    "..jpg",
			Expected: "._lowQuality.jpg",
		},
		{
			Title:    "Very long filename",
			Input:    "very_long_filename_that_might_cause_issues_in_some_scenarios_but_should_still_work_correctly.jpg",
			Expected: "very_long_filename_that_might_cause_issues_in_some_scenarios_but_should_still_work_correctly_lowQuality.jpg",
		},
		{
			Title:    "Filename with numbers",
			Input:    "image123.jpg",
			Expected: "image123_lowQuality.jpg",
		},
		{
			Title:    "Filename with mixed case extension",
			Input:    "image.JPG",
			Expected: "image_lowQuality.JPG",
		},
		{
			Title:    "Filename with uppercase extension",
			Input:    "PHOTO.PNG",
			Expected: "PHOTO_lowQuality.PNG",
		},
		{
			Title:    "File without extension",
			Input:    "noextension",
			Expected: "noextension_lowQuality",
		},
		{
			Title:    "File with path but no extension",
			Input:    "/path/to/noextension",
			Expected: "noextension_lowQuality",
		},
		{
			Title:    "Just an extension",
			Input:    ".jpg",
			Expected: "_lowQuality.jpg",
		},
		{
			Title:    "Filename with parentheses",
			Input:    "image(1).jpg",
			Expected: "image(1)_lowQuality.jpg",
		},
		{
			Title:    "Filename with brackets",
			Input:    "image[copy].jpg",
			Expected: "image[copy]_lowQuality.jpg",
		},
		{
			Title:    "Filename with single quote",
			Input:    "user's_photo.jpg",
			Expected: "user's_photo_lowQuality.jpg",
		},
		{
			Title:    "Filename with double quote",
			Input:    "image\"test\".jpg",
			Expected: "image\"test\"_lowQuality.jpg",
		},
		{
			Title:    "Very short filename",
			Input:    "a.jpg",
			Expected: "a_lowQuality.jpg",
		},
		{
			Title:    "Single character filename",
			Input:    "x",
			Expected: "x_lowQuality",
		},
		{
			Title:    "Filename with percent encoding characters",
			Input:    "image%20with%20spaces.jpg",
			Expected: "image%20with%20spaces_lowQuality.jpg",
		},
		{
			Title:    "Filename with ampersand",
			Input:    "cats&dogs.jpg",
			Expected: "cats&dogs_lowQuality.jpg",
		},
		{
			Title:    "Filename with plus sign",
			Input:    "image+copy.jpg",
			Expected: "image+copy_lowQuality.jpg",
		},
		{
			Title:    "Filename with equals sign",
			Input:    "formula=result.jpg",
			Expected: "formula=result_lowQuality.jpg",
		},
		{
			Title:    "Filename with hash",
			Input:    "image#1.jpg",
			Expected: "image#1_lowQuality.jpg",
		},
		{
			Title:    "Windows reserved filename with extension",
			Input:    "CON.jpg",
			Expected: "CON_lowQuality.jpg",
		},
		{
			Title:    "Path with mixed separators",
			Input:    "path\\to/mixed\\separators/image.jpg",
			Expected: "image_lowQuality.jpg",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Title, func(t *testing.T) {
			result := GetLowQualityImageName(testCase.Input)
			if result != testCase.Expected {
				t.Errorf("Expected %q, but got %q for input %q",
					testCase.Expected, result, testCase.Input)
			}
		})
	}
}
