package enums

type FileType string

const (
	VIDEO    FileType = "video"
	AUDIO    FileType = "audio"
	SUBTITLE FileType = "subtitle"
	IMAGE    FileType = "image"
	UNKNOWN  FileType = "unknown"
)
