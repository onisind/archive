package models

import (
	"time"
)

type Document struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	Filename  string    `json:"filename"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"-"`
	MongoID   []string  `gorm:"type:text[]" json:"mongo_ids"`
}
