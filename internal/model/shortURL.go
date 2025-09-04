package model

import "gorm.io/gorm"

type ShortURL struct {
	gorm.Model
	URL         string `gorm:"not null"`
	ShortCode   string `gorm:"uniqueIndex; size:20; not null"`
	AccessCount int
}
