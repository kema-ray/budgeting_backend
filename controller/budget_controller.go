package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kema-ray/home-budgeting-app/config"
	"github.com/kema-ray/home-budgeting-app/models"
)

func AddBudget(c *gin.Context) {
	userID, _ := c.Get("userID")

	var budget models.Budget
	if err := c.ShouldBindJSON(&budget); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	budget.UserID = userID.(uint)

	// Save the budget to the database
	if err := config.DB.Create(&budget).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add budget"})
		return
	}

	// Respond with the added budget data
	c.JSON(http.StatusOK, gin.H{
		"message": "Budget added successfully",
		"data":    budget,  // Return the added budget data
	})
}

func GetAllBudgets(c *gin.Context) {
	var budgets []models.Budget

	// Query the database to get all budgets
	if err := config.DB.Find(&budgets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve budgets"})
		return
	}

	// Respond with the list of budgets
	c.JSON(http.StatusOK, gin.H{
		"message": "Budgets retrieved successfully",
		"data":    budgets,
	})
}

func GetUserBudgets(c *gin.Context) {
	// Get the logged-in user's ID from the context
	userID, _ := c.Get("userID")

	var budgets []models.Budget

	// Query the database for budgets that belong to the logged-in user
	if err := config.DB.Where("user_id = ?", userID).Find(&budgets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve budgets for the user"})
		return
	}

	// Respond with the list of budgets for the logged-in user
	c.JSON(http.StatusOK, gin.H{
		"message": "User budgets retrieved successfully",
		"data":    budgets,
	})
}


