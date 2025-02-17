package models

import (
	"gorm.io/gorm"
	"time"
)

type Expense struct {
	gorm.Model
	UserID   uint      `gorm:"not null" json:"user_id"`
	Category string    `gorm:"not null" json:"category"`
	Amount   float64   `gorm:"not null" json:"amount"`
	Date     time.Time `gorm:"not null" json:"date"`
	User     User      `gorm:"foreignKey:UserID"`
}
