package models

import (
	"backend/pkg/enums"
)

type FileTreeDto struct {
	Name                  string        `json:"name"`
	Id                    string        `json:"id"`
	ThumbnailId           string        `json:"thumbnailId"`
	LowQualityThumbnailId string        `json:"lowQualityThumbnailId"`
	Files                 []FileDto     `json:"files"`
	Children              []FileTreeDto `json:"children"`
}

type FileDto struct {
	Id                    string         `json:"id"`
	Name                  string         `json:"name"`
	Type                  enums.FileType `json:"fileType"`
	AssociatedAudioFileId string         `json:"audioFileId"`
	LowQualityImageId     string         `json:"lowQualityImageId"`
}
