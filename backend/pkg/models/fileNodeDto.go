package models

import "backend/pkg/enums/fileType"

type FileNodeDto struct {
	Id                    string            `json:"id"`
	Name                  string            `json:"name"`
	Type                  fileType.FileType `json:"fileType"`
	AssociatedAudioFileId string            `json:"audioFileId"`
	LowQualityImageId     string            `json:"lowQualityImageId"`
}
