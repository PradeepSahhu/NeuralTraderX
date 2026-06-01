package main

import (
	"github.com/PradeepSahhu/NeuralTraderX/api"
	"github.com/PradeepSahhu/NeuralTraderX/internal/database"
	"github.com/PradeepSahhu/NeuralTraderX/internal/handler"
	"github.com/PradeepSahhu/NeuralTraderX/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	api.RegisterHealthApiRoutes(router)

	// database GORM
	db := database.RegisterPostgresSql()
	tempRepo := repository.NewTempRepository(db)
	// if err := tempRepo.CreateTemp("sandeep", "singh"); err != nil {
	// 	log.Fatalf("failed to create temp row: %v", err)
	// }
	tempHandler := handler.NewTempHandler(tempRepo)

	api.RegisterBasicRoutes(router, tempHandler)

	router.Run(":3000")
}
