package utilities

import (
	"backend/pkg/constants"
	"backend/pkg/enums"
	"path"
	"slices"
)

func GetFileType(fileName string) enums.FileType {

	extension := path.Ext(fileName)

	if slices.Contains(constants.AllowedAudioFileExtensions, extension) {
		return enums.AUDIO
	}

	if slices.Contains(constants.AllowedVideoFileExtensions, extension) {
		return enums.VIDEO
	}

	if slices.Contains(constants.AllowedImageFileExtensions, extension) {
		return enums.IMAGE
	}

	if slices.Contains(constants.AllowedSubtitleFileExtensions, extension) {
		return enums.SUBTITLE
	}

	return enums.UNKNOWN
}
