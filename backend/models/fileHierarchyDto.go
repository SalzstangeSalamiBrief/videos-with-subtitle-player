package models

import "github.com/google/uuid"

type FileHierarchyDto struct {
	AudioFiles AudioFileDictionary
	Children   map[string]FileHierarchyDto
}

type AudioFileDictionary map[string][]FileItem // TODO LATER CHANGE TO TUPLE

type FileItem struct {
	Id   uuid.UUID
	Name string
}
