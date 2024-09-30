package lib

import (
	"backend/enums"
	"path"
)

func GetFileType(fileName string) enums.FileType {

	extension := path.Ext(fileName)

	switch extension {
	case ".mp4":
		return enums.VIDEO
	case ".vtt":
		return enums.SUBTITLE
	case ".mp3":
		return enums.AUDIO
	case ".wav":
		return enums.AUDIO
	case ".webp":
		return enums.IMAGE
	case ".avif":
		return enums.IMAGE
	case ".png":
		return enums.IMAGE
	case ".jpeg":
		return enums.IMAGE
	case ".jpg":
		return enums.IMAGE
	default:
		return enums.UNKNOWN
	}
}
