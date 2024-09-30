package logic

import (
	"backend/enums"
	"backend/lib/utilities/models"
	"path"
	"slices"
)

func GetFileType(fileName string) enums.FileType {

	extension := path.Ext(fileName)

	if slices.Contains(models.AllowedAudioFileExtensions, extension) {
		return enums.AUDIO
	}

	if slices.Contains(models.AllowedVideoFileExtensions, extension) {
		return enums.VIDEO
	}

	if slices.Contains(models.AllowedImageFileExtensions, extension) {
		return enums.IMAGE
	}

	if slices.Contains(models.AllowedSubtitleFileExtensions, extension) {
		return enums.SUBTITLE
	}

	return enums.UNKNOWN
}
