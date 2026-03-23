package models

import (
	"time"

	"github.com/google/uuid"
)

type FolderNode struct {
	ID                    uint `gorm:"primary_key"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	FolderId              string       `gorm:"type:UUID;not null;unique"`
	Name                  string       `gorm:"unique"`
	Path                  string       `gorm:"unique"`
	ThumbnailId           string       `json:"thumbnailId"`
	LowQualityThumbnailId string       `json:"lowQualityThumbnailId"`
	Files                 []FileNode   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ParentFolderId;references:FolderId;"`
	ParentFolderId        *string      `gorm:"type:UUID;index"`
	ParentFolder          *FolderNode  `gorm:"foreignKey:ParentFolderId;references:FolderId"`
	ChildFolders          []FolderNode `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:ParentFolderId;references:FolderId"`
	Tags                  []Tag        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;many2many:folder_node_to_tags;"`
}

func (node *FolderNode) ToDto() FolderNodeDto {
	dto := FolderNodeDto{
		Id:                    node.FolderId,
		Name:                  node.Name,
		ThumbnailId:           node.ThumbnailId,
		LowQualityThumbnailId: node.LowQualityThumbnailId,
		Files:                 &[]FileNodeDto{},
		Children:              &[]FolderNodeDto{},
	}
	// TODO DOES NOT WORK PROPERLY
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

func NodesToSingleTree(nodes []FolderNode) FolderNodeDto {
	root := FolderNodeDto{
		Name:                  "",
		Id:                    uuid.New().String(),
		ThumbnailId:           "",
		LowQualityThumbnailId: "",
		Files:                 nil,
		Children:              nil,
	}

	if nodes != nil {
		childNodes := make([]FolderNodeDto, len(nodes))
		for i, child := range nodes {
			childDto := child.ToDto()
			childNodes[i] = childDto
		}

		root.Children = &childNodes
	}

	return root
}

func NodesToFlatHierarchy(nodes []FolderNode) []FolderNode {
	flatHierarchy := make([]FolderNode, 0)
	for _, node := range nodes {
		flatHierarchy = append(flatHierarchy, node.transformTreeToFlatHierarchy()...)
	}

	return flatHierarchy
}

func (node FolderNode) transformTreeToFlatHierarchy() []FolderNode {
	// copy children & clear children before copying the node into folder to prevent duplicates
	children := node.ChildFolders
	node.ChildFolders = make([]FolderNode, 0)

	flatHierarchy := []FolderNode{node}
	for _, child := range children {
		flatHierarchy = append(flatHierarchy, child.transformTreeToFlatHierarchy()...)
	}

	return flatHierarchy
}
