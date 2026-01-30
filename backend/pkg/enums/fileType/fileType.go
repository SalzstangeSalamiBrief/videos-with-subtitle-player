package fileType

type FileType string

const (
	VIDEO    FileType = "Video"
	AUDIO    FileType = "Audio"
	SUBTITLE FileType = "Subtitle"
	IMAGE    FileType = "Image"
	UNKNOWN  FileType = "Unknown"
)
