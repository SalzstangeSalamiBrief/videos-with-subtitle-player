package models

import (
	"time"
)

type Tag struct {
	// Dont use gorm.Model to prevent soft delete
	ID            uint `gorm:"primary_key"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string         `gorm:"unique"`
	FileTreeItems []FileTreeItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;many2many:file_tree_item_to_tags;"`
}
