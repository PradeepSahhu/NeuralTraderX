package api

import (
	"github.com/PradeepSahhu/NeuralTraderX/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterHealthApiRoutes(router *gin.Engine) {
	router.GET("/health", handler.HealthCheck)
}

func RegisterBasicRoutes(router *gin.Engine, tempHandler *handler.TempHandler) {
	router.GET("/status", handler.HealthCheck)
	router.POST("/create", tempHandler.CreateTempHandler)
}
