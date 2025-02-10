package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kema-ray/home-budgeting-app/config"
	"github.com/kema-ray/home-budgeting-app/routes"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize PSQL database
	config.ConnectDatabase()

	// Initialize the Gin Router
	router := gin.Default()

	// Set up all routes
	routes.SetupRoutes(router)

	// Get the port from environment variables or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("🚀 Server is running at http://localhost:" + port)
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}