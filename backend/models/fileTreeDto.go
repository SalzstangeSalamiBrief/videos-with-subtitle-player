package models

import "github.com/google/uuid"

// TODO maybe swagger??
type FileTreeDto struct {
	Name       string `json:"name"`
	Id         uuid.UUID `json:"id"`
	AudioFiles []AudioFileDto `json:"audioFiles"`
	Children   []FileTreeDto `json:"children"`
}

type AudioFileDto struct {
	Name         string `json:"name"`
	SubtitleFile FileItemDto `json:"subtitleFile"`
	AudioFile    FileItemDto `json:"audioFile"`
}

type FileItemDto struct {
	Id   uuid.UUID `json:"id"`
	Name string `json:"name"`
}
