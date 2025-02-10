package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model

	UserID uint `gorm:"not null"`
	BudgetID uint `gorm:"not null"`
	Amount float64 `gorm:"not null"`
}