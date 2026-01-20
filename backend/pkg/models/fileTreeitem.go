package models

import (
	"backend/pkg/enums"
)

type FileTreeItem struct {
	Id                    string
	Path                  string
	Name                  string
	Type                  enums.FileType
	AssociatedAudioFileId *string
	LowQualityImageId     *string
}

func (item *FileTreeItem) ToFileDto() FileDto {
	dto := FileDto{
		Id:   item.Id,
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
