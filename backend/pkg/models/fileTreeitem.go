package models

import (
	"backend/pkg/enums"
)

type FileTreeItem struct {
	Id                    string
	Path                  string
	Name                  string
	Type                  enums.FileType
	AssociatedAudioFileId string
	LowQualityImageId     string
}
