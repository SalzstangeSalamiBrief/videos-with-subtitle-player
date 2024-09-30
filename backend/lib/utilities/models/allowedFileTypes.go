package models

import "slices"

var AllowedAudioFileExtensions = []string{".mp3", ".wav"}

var AllowedVideoFileExtensions = []string{".mp4"}

var AllowedContinuousFileExtensions = slices.Concat(AllowedAudioFileExtensions, AllowedVideoFileExtensions)

var AllowedImageFileExtensions = []string{".avif", ".png", ".jpeg", ".jpg", ".webp"}

var AllowedSubtitleFileExtensions = []string{".vtt"}

var AllowedDiscreteFileExtensions = slices.Concat(AllowedImageFileExtensions, AllowedSubtitleFileExtensions)
