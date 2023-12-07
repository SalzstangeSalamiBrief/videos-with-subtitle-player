package models

import "github.com/google/uuid"

type FileTreeItem struct {
	Id   uuid.UUID
	Path string
	Name string
}
