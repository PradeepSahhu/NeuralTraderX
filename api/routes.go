package api

import (
	"github.com/PradeepSahhu/NeuralTraderX/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(router *gin.Engine) {
	router.GET("/health", handler.HealthCheck)
}
