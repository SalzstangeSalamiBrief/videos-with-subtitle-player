package models

import "time"

type Tag struct {
	Id        uint   `gorm:"primary_key"`
	Name      string `gorm:"unique"`
	CreatedAt time.Time
}
