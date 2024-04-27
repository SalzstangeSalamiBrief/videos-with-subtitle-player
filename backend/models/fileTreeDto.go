package models

import (
	"backend/enums"
)

type FileTreeDto struct {
	Name     string        `json:"name"`
	Id       string        `json:"id"`
	Files    []FileDto     `json:"files"`
	Children []FileTreeDto `json:"children"`
}

type FileDto struct {
	Id                    string         `json:"id"`
	Name                  string         `json:"name"`
	Type                  enums.FileType `json:"fileType"`
	AssociatedAudioFileId string         `json:"audioFileId"`
}
