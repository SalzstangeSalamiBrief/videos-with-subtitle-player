package models

import "backend/models/enums"

type FileTreeItem struct {
	Id   string
	Path string
	Name string
	Type enums.FileType
	// used to associate subtilte files with audio files
	AssociatedAudioFileId string
}
