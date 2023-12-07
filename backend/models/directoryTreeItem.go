package models

import "github.com/google/uuid"

type DirectoryTreeItem struct {
	Id           uuid.UUID
	Path         string
	Name         string
	Children     []DirectoryTreeItem
	AudioFile    FileTreeItem
	SubtitleFile FileTreeItem
}
