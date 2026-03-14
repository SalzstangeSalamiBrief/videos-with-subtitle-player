package models

import "time"

type FolderNode struct {
	ID                    uint `gorm:"primary_key"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	FolderId              string       `gorm:"type:UUID;not null;unique"`
	Name                  string       `gorm:"unique"`
	Path                  string       `gorm:"unique"`
	ThumbnailId           string       `json:"thumbnailId"`
	LowQualityThumbnailId string       `json:"lowQualityThumbnailId"`
	Files                 []FileNode   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:ParentFolderId;references:FolderId;"`
	ParentFolderId        *string      `gorm:"type:UUID;index"`
	ParentFolder          *FolderNode  `gorm:"foreignKey:ParentFolderId;references:FolderId"`
	ChildFolders          []FolderNode `gorm:"constraint:OnUpdate:CASCADE;foreignKey:ParentFolderId;references:FolderId"`
}

func (node *FolderNode) ToDto() FolderNodeDto {
	dto := FolderNodeDto{
		Id:                    node.FolderId,
		Name:                  node.Name,
		ThumbnailId:           node.ThumbnailId,
		LowQualityThumbnailId: node.LowQualityThumbnailId,
		Files:                 nil,
		Children:              nil,
	}

	if node.ChildFolders != nil {
		childFolders := make([]FolderNodeDto, len(node.ChildFolders))
		for i, child := range node.ChildFolders {
			childDto := child.ToDto()
			childFolders[i] = childDto
		}

		dto.Children = &childFolders
	}

	if node.Files != nil {
		files := make([]FileNodeDto, len(node.Files))
		for i, file := range node.Files {
			fileDto := file.ToDto()
			files[i] = fileDto
		}

		dto.Files = &files
	}

	return dto
}
