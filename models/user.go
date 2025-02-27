package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	Budgets  []Budget  `gorm:"foreignKey:UserID"`
	Expenses []Expense `gorm:"foreignKey:UserID"`
}

type RegisterUserInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}
