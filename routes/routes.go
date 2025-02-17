package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kema-ray/home-budgeting-app/controller"
	"github.com/kema-ray/home-budgeting-app/middleware"
)

func SetupRoutes(router *gin.Engine) {
	// public routes
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", controller.Register)
		authRoutes.POST("/login", controller.Login)
	}

	// Protected routes(require JWT Token)
	apiRoutes := router.Group("/api")
	apiRoutes.Use(middleware.AuthMiddleware()) // Apply AuthMiddleware to all routes in apiRoutes
	{
		// Budget Routes
		apiRoutes.POST("/budgets", controller.AddBudget)
		apiRoutes.GET("/list-budgets", controller.GetAllBudgets) // Get all budgets
		apiRoutes.GET("/user/budgets", controller.GetUserBudgets) // Get budgets for the logged-in user


		// Expense Routes
		// apiRoutes.GET("/expenses", controller.GetExpenses)
	}

}