package lib

import (
	"backend/enums"
	"path"
)

func GetFileType(fileName string) enums.FileType {
	extension := path.Ext(fileName)
	switch extension {
	case ".mp3":
		return enums.AUDIO
	case ".mp4":
		return enums.VIDEO
	case ".wav":
		return enums.AUDIO
	case ".vtt":
		return enums.SUBTITLE
	default:
		return enums.UNKNOWN
	}
}
