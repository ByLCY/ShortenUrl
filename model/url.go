package model

import (
	"time"
)

type Url struct {
	Id        uint64 `gorm:"autoIncrement,primaryKey"`
	LongURL   string `gorm:"uniqueIndex"`
	ShortURL  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Url) TableName() string {
	return "urls"
}
