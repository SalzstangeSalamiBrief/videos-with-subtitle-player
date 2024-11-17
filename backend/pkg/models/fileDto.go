package models

import "backend/pkg/enums"

type FileDto struct {
	Id                    string         `json:"id"`
	Name                  string         `json:"name"`
	Type                  enums.FileType `json:"fileType"`
	AssociatedAudioFileId string         `json:"audioFileId"`
}
