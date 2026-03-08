package models

import (
	"time"
)

type Tag struct {
	// Dont use gorm.Model to prevent soft delete
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"unique"`
}
