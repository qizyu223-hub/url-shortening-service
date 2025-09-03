package model

import "gorm.io/gorm"

type ShortURL struct {
	gorm.Model
	URL         string `gorm:"not null"`
	ShortCode   string `gorm:"unique"`
	AccessCount int
}
