package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	Budgets  []Budget  `gorm:"foreignKey:UserID"`
	Expenses []Expense `gorm:"foreignKey:UserID"`
}
