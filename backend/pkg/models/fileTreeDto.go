package models

type FileTreeDto struct {
	Name        string        `json:"name"`
	Id          string        `json:"id"`
	ThumbnailId string        `json:"thumbnailId"`
	Files       []FileDto     `json:"files"`
	Children    []FileTreeDto `json:"children"`
}

func (tree *FileTreeDto) GetIndexOfChildByName(name string) int {
	for i, child := range tree.Children {
		if child.Name == name {
			return i
		}
	}

	return -1
}
