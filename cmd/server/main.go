package main

import (
	"github.com/PradeepSahhu/NeuralTraderX/api"
	"github.com/PradeepSahhu/NeuralTraderX/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	api.RegisterHealthApiRoutes(router)
	api.RegisterBasicRoutes(router)
	database.RegisterPostgresSql()

	router.Run(":3000")
}
