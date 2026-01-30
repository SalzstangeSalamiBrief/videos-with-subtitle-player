package models

import (
	"backend/pkg/enums"
	"time"
)

type FileTreeItem struct {
	// TODO DO I NEED FILE_ID IF I HAVE gorm.Model.ID?
	// Dont use gorm.Model to prevent soft delete
	ID                    uint `gorm:"primary_key"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	FileId                string `gorm:"type:UUID"`
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
