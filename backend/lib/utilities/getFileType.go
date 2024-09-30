package lib

import (
	"backend/enums"
	"path"
	"slices"
)

func GetFileType(fileName string) enums.FileType {

	extension := path.Ext(fileName)

	if slices.Contains(AllowedAudioFileExtensions, extension) {
		return enums.AUDIO
	}

	if slices.Contains(AllowedVideoFileExtensions, extension) {
		return enums.VIDEO
	}

	if slices.Contains(AllowedImageFileExtensions, extension) {
		return enums.IMAGE
	}

	if slices.Contains(AllowedSubtitleFileExtensions, extension) {
		return enums.SUBTITLE
	}

	return enums.UNKNOWN
}
