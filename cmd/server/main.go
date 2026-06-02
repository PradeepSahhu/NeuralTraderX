package main

import (
	"github.com/PradeepSahhu/NeuralTraderX/api"
	"github.com/PradeepSahhu/NeuralTraderX/internal/database"
	"github.com/PradeepSahhu/NeuralTraderX/internal/handler"
	"github.com/PradeepSahhu/NeuralTraderX/internal/repository"
	"github.com/gin-gonic/gin"
)

// need : repository need -> handler -> routes

// db -> repository -> handler -> routes register karna hai
func main() {

	router := gin.Default()

	// database GORM
	db := database.RegisterPostgresSql()

	// repo creation
	tempRepo := repository.NewTempRepository(db)

	// handler creation
	tempHandler := handler.NewTempHandler(tempRepo)

	registrars := []api.RouteRegistrar{
		tempHandler,
	}

	apiGroup := router.Group("/api/v1")

	for _, r := range registrars {
		r.RegisterRoutes(apiGroup)
	}

	router.Run(":3000")
}
