package models

import (
	"time"

	"github.com/lib/pq"
)

type Document struct {
	ID        uint           `gorm:"primaryKey" json:"-"`
	Filename  string         `json:"filename"`
	Author    string         `json:"author"`
	CreatedAt time.Time      `json:"-"`
	MongoIDs  pq.StringArray `gorm:"type:text[]" json:"mongo_ids"`
}
