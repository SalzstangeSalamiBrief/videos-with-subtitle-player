package models

import (
	"backend/pkg/enums/fileType"
	"time"
)

type FileNode struct {
	// TODO DO I NEED FILE_ID IF I HAVE gorm.Model.ID?
	// Dont use gorm.Model to prevent soft delete
	ID                    uint `gorm:"primary_key"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	FileId                string `gorm:"type:UUID"`
	Path                  string
	Name                  string
	Type                  fileType.FileType `gorm:"type:file_type;not null"`
	AssociatedAudioFileId *string
	LowQualityImageId     *string
	ParentFolderId        string `gorm:"type:UUID;index"`
}

func (node *FileNode) ToDto() FileNodeDto {
	dto := FileNodeDto{
		Id:   node.FileId,
		Name: node.Name,
		Type: node.Type,
	}

	if node.AssociatedAudioFileId != nil {
		dto.AssociatedAudioFileId = *node.AssociatedAudioFileId
	}

	if node.LowQualityImageId != nil {
		dto.LowQualityImageId = *node.LowQualityImageId
	}

	return dto
}
