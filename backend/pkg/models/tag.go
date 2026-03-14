package models

import (
	"time"
)

type Tag struct {
	// Dont use gorm.Model to prevent soft delete
	ID            uint `gorm:"primary_key"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string     `gorm:"unique"`
	FileTreeItems []FileNode `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;many2many:file_node_to_tags;"`
}
