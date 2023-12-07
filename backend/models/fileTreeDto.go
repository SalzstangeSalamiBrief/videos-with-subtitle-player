package models

import "github.com/google/uuid"

type FileTreeDto struct {
	AudioFiles AudioFileDictionary
	Children   map[string]FileTreeDto
}

type AudioFileDictionary map[string][]FileItem

type FileItem struct {
	Id   uuid.UUID
	Name string
}
