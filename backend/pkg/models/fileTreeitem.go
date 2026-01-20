package models

import (
	"backend/pkg/enums"
	"gorm.io/gorm"
)

type FileTreeItem struct {
	gorm.Model
	FileId                string
	Path                  string
	Name                  string
	Type                  enums.FileType `gorm:"type:file_type;not null"`
	AssociatedAudioFileId *string
	LowQualityImageId     *string
}

func (item *FileTreeItem) ToFileDto() FileDto {
	dto := FileDto{
		Id:   item.FileId,
		Name: item.Name,
		Type: item.Type,
	}

	if item.AssociatedAudioFileId != nil {
		dto.AssociatedAudioFileId = *item.AssociatedAudioFileId
	}

	if item.LowQualityImageId != nil {
		dto.LowQualityImageId = *item.LowQualityImageId
	}

	return dto
}
