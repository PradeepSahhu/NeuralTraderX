package main

import (
	"github.com/PradeepSahhu/NeuralTraderX/api"
	"github.com/PradeepSahhu/NeuralTraderX/internal/database"
	"github.com/PradeepSahhu/NeuralTraderX/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	api.RegisterHealthApiRoutes(router)
	api.RegisterBasicRoutes(router)

	// database GORM
	db := database.RegisterPostgresSql()
	repository.NewTempRepository(db)

	router.Run(":3000")
}
