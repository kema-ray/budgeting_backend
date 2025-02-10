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
	}

	budget.UserID = userID.(uint)

	if err := config.DB.Create(&budget).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add budget"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Budget added successfully"})
}