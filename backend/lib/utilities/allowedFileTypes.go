package lib

import "slices"

// TODO CREATE A FOLDER FOR THIS
var AllowedAudioFileExtensions = []string{".mp3", ".wav"}

var AllowedVideoFileExtensions = []string{".mp4"}

var AllowedContinuousFileExtensions = slices.Concat(AllowedAudioFileExtensions, AllowedVideoFileExtensions)

var AllowedImageFileExtensions = []string{".avif", ".png", ".jpeg", ".jpg"}

var AllowedSubtitleFileExtensions = []string{".vtt"}

var AllowedDiscreteFileExtensions = slices.Concat(AllowedImageFileExtensions, AllowedSubtitleFileExtensions)
