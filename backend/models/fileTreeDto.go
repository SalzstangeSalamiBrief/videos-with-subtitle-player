package models

import "github.com/google/uuid"

// TODO maybe swagger??
type FileTreeDto struct {
	Name       string
	Id         uuid.UUID
	AudioFiles []AudioFileDto
	Children   []FileTreeDto
}

type AudioFileDto struct {
	Name         string
	SubtitleFile FileItemDto
	AudioFile    FileItemDto
}

type FileItemDto struct {
	Id   uuid.UUID
	Name string
}
