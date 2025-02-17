package models

import "gorm.io/gorm"

type Budget struct {
	gorm.Model
	UserID       uint    `gorm:"not null" json:"user_id"`
	Category     string  `gorm:"not null" json:"category"`
	Limit        float64 `gorm:"not null" json:"limit"`
	CurrentSpent float64 `gorm:"default:0" json:"current_spent"`
	User         User    `gorm:"foreignKey:UserID"`
}
