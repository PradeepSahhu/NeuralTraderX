package handler

import (
	"github.com/PradeepSahhu/NeuralTraderX/internal/repository"
)

type HealthHandler struct {
	repo repository.TempRepository
}

// func HealthCheck(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"Status": "Working",
// 	})
// }
