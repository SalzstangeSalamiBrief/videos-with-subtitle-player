package utilities

import (
	"backend/pkg/constants"
	"backend/pkg/enums/fileType"
	"path"
	"slices"
)

func GetFileType(fileName string) fileType.FileType {

	extension := path.Ext(fileName)

	if slices.Contains(constants.AllowedAudioFileExtensions, extension) {
		return fileType.AUDIO
	}

	if slices.Contains(constants.AllowedVideoFileExtensions, extension) {
		return fileType.VIDEO
	}

	if slices.Contains(constants.AllowedImageFileExtensions, extension) {
		return fileType.IMAGE
	}

	if slices.Contains(constants.AllowedSubtitleFileExtensions, extension) {
		return fileType.SUBTITLE
	}

	return fileType.UNKNOWN
}
