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

/**
 * CORSMiddleware handles Cross-Origin Resource Sharing issues that occur 
 * when the frontend (running on a different origin) attempts to make requests 
 * to the backend API.
 *
 * This middleware:
 * 1. Sets headers to allow cross-origin requests from any origin (*)
 * 2. Allows credentials to be included in cross-origin requests
 * 3. Specifies which HTTP headers can be used in requests
 * 4. Defines which HTTP methods are permitted (GET, POST, PUT, DELETE, OPTIONS)
 * 5. Handles preflight OPTIONS requests by returning a 204 No Content status
 *
 * Without this middleware, browsers would block requests from your frontend
 * to your backend due to the Same-Origin Policy security feature.
 */
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Allow requests from any origin (using * is convenient for development
		// but should be restricted to specific origins in production)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		
		// Allow cookies and authorization headers to be sent with requests
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		
		// Define which headers the client is allowed to send in requests
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		
		// Define which HTTP methods are permitted
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		// Handle preflight requests (OPTIONS method)
		// Browsers send an OPTIONS request before certain cross-origin requests
		// to check if the actual request is safe to send
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // Respond with 204 No Content status
			return
		}

		// Continue to the next middleware or route handler
		c.Next()
	}
}

/**
 * The main function is the entry point of the application.
 * It performs the following tasks:
 * 1. Loads environment variables
 * 2. Initializes the database connection
 * 3. Sets up the Gin HTTP router with middleware
 * 4. Configures API routes
 * 5. Starts the web server on the specified port
 */
func main() {
	// Load environment variables from .env file
	// This allows configuration to be stored outside the codebase
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize PostgreSQL database connection
	// The connection details are specified in environment variables
	config.ConnectDatabase()

	// Initialize the Gin Router with default middleware
	// (logger and recovery middleware)
	router := gin.Default()

	// Apply the CORS middleware to all routes
	// This ensures all endpoints can be accessed from the frontend
	router.Use(CORSMiddleware())

	// Set up all application routes
	// Routes are defined in the routes package
	routes.SetupRoutes(router)

	// Get the port from environment variables or default to 8080
	// This allows the port to be configured via environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Print a startup message to the console
	fmt.Println("🚀 Server is running at http://localhost:" + port)
	
	// Start the HTTP server on the specified port
	// This is a blocking call that runs until the server is stopped
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}