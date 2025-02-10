package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kema-ray/home-budgeting-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the global database connection
var DB *gorm.DB

// ConnectDatabase initializes the database connection
func ConnectDatabase () {
	_= godotenv.Load() // Load .env file

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database!")
	}

	fmt.Println("📌 Database connected successfully!")
	DB.AutoMigrate(&models.User{}, &models.Budget{})
}