package models

import "time"

type Migration struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	Checksum  string
	Name      string
}

func (Migration) TableName() string {
	return "migrations"
}
