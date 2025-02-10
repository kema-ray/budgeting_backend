package models

import "gorm.io/gorm"

type Budget struct {
	gorm.Model

	UserID uint `gorm:"not null"`
	Category string `gorm:"not null"`
	Amount float64 `gorm:"not null"`
}