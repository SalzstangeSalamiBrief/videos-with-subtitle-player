package models

type FolderNodeDto struct {
	Name                  string           `json:"name"`
	Id                    string           `json:"id"`
	ThumbnailId           string           `json:"thumbnailId"`
	LowQualityThumbnailId string           `json:"lowQualityThumbnailId"`
	Files                 *[]FileNodeDto   `json:"files"`
	Children              *[]FolderNodeDto `json:"children"`
}
