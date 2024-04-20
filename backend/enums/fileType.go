package enums

type FileType string

const (
	VIDEO    FileType = "video"
	AUDIO    FileType = "audio"
	SUBTITLE FileType = "subtitle"
	UNKNOWN  FileType = "unknown"
)
