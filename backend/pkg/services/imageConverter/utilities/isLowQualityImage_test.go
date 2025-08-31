package utilities

import (
	"backend/pkg/models"
	"testing"
)

func TestIsLowQualityImage(t *testing.T) {
	testCases := []models.TestData[string, bool]{
		{
			Title:    "Valid low quality image",
			Input:    "image_lowQuality.webp",
			Expected: true,
		},
		{
			Title:    "Valid low quality image with path",
			Input:    "folder/image_lowQuality.webp",
			Expected: true,
		},
		{
			Title:    "Valid low quality image with deep path",
			Input:    "assets/images/gallery/photo_lowQuality.webp",
			Expected: true,
		},
		{
			Title:    "Valid low quality image with absolute path",
			Input:    "/home/user/images/picture_lowQuality.webp",
			Expected: true,
		},
		{
			Title:    "Valid low quality image with Windows path",
			Input:    "C:\\Users\\user\\images\\photo_lowQuality.webp",
			Expected: true,
		},
		{
			Title:    "Regular image - not low quality",
			Input:    "image.jpg",
			Expected: false,
		},
		{
			Title:    "Regular webp image - not low quality",
			Input:    "image.webp",
			Expected: false,
		},
		{
			Title:    "Image with low quality suffix but wrong extension",
			Input:    "image_lowQuality.jpg",
			Expected: false,
		},
		{
			Title:    "Image with low quality suffix but wrong extension - png",
			Input:    "image_lowQuality.png",
			Expected: false,
		},
		{
			Title:    "Image with partial suffix match",
			Input:    "image_low.webp",
			Expected: false,
		},
		{
			Title:    "Image with quality in name but not suffix",
			Input:    "quality_image.webp",
			Expected: false,
		},
		{
			Title:    "Empty string",
			Input:    "",
			Expected: false,
		},
		{
			Title:    "Just the extension",
			Input:    ".webp",
			Expected: false,
		},
		{
			Title:    "Just the suffix",
			Input:    "_lowQuality",
			Expected: false,
		},
		{
			Title:    "Filename with spaces",
			Input:    "my photo_lowQuality.webp",
			Expected: true,
		},
		{
			Title:    "Filename with special characters",
			Input:    "image-test_01_lowQuality.webp",
			Expected: true,
		},
		{
			Title:    "Filename with Unicode characters",
			Input:    "测试图片_lowQuality.webp",
			Expected: true,
		},
		{
			Title:    "Filename with umlauts",
			Input:    "bild_äöü_lowQuality.webp",
			Expected: true,
		},
		{
			Title:    "Hidden file",
			Input:    ".hidden_lowQuality.webp",
			Expected: true,
		},
		{
			Title:    "Very long filename",
			Input:    "very_long_filename_that_might_cause_issues_in_some_scenarios_lowQuality.webp",
			Expected: true,
		},
		{
			Title:    "Filename with numbers",
			Input:    "image123_lowQuality.webp",
			Expected: true,
		},
		{
			Title:    "Uppercase extension",
			Input:    "image_lowQuality.WEBP",
			Expected: false,
		},
		{
			Title:    "Mixed case extension",
			Input:    "image_lowQuality.Webp",
			Expected: false,
		},
		{
			Title:    "Suffix in middle of filename",
			Input:    "image_lowQuality.webp_backup.webp",
			Expected: true,
		},
		{
			Title:    "Multiple occurrences of pattern",
			Input:    "image_lowQuality.webp_lowQuality.webp",
			Expected: true,
		},
		{
			Title:    "Pattern at beginning",
			Input:    "_lowQuality.webp_image.jpg",
			Expected: true,
		},
		{
			Title:    "Filename with parentheses",
			Input:    "image(1)_lowQuality.webp",
			Expected: true,
		},
		{
			Title:    "Filename with brackets",
			Input:    "image[copy]_lowQuality.webp",
			Expected: true,
		},
		{
			Title:    "Very short filename",
			Input:    "a_lowQuality.webp",
			Expected: true,
		},
		{
			Title:    "Path with mixed separators",
			Input:    "path\\to/image_lowQuality.webp",
			Expected: true,
		},
		{
			Title:    "Relative path with dots",
			Input:    "../images/./photo_lowQuality.webp",
			Expected: true,
		},
		{
			Title:    "Similar but incorrect pattern - missing underscore",
			Input:    "imagelowQuality.webp",
			Expected: false,
		},
		{
			Title:    "Similar but incorrect pattern - wrong case",
			Input:    "image_LowQuality.webp",
			Expected: false,
		},
		{
			Title:    "Similar but incorrect pattern - extra characters",
			Input:    "image_lowQualityx.webp",
			Expected: false,
		},
		{
			Title:    "File without extension",
			Input:    "image_lowQuality",
			Expected: false,
		},
		{
			Title:    "Double extension with correct pattern",
			Input:    "archive_lowQuality.webp.backup",
			Expected: true,
		},
		{
			Title:    "Pattern in directory name, not filename",
			Input:    "folder_lowQuality.webp/image.jpg",
			Expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Title, func(t *testing.T) {
			result := IsLowQualityImage(testCase.Input)
			if result != testCase.Expected {
				t.Errorf("Expected %v, but got %v for input %q",
					testCase.Expected, result, testCase.Input)
			}
		})
	}
}
