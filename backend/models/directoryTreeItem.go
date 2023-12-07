package models

import "github.com/google/uuid"

// TODO ARRAY INSTEAD OF NESTED STRUCT
type DirectoryTreeItem struct {
	Id           uuid.UUID
	Path         string
	Name         string
	Children     []DirectoryTreeItem
	AudioFile    FileTreeItem
	SubtitleFile FileTreeItem
}
