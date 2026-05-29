package main

import (
	"github.com/PradeepSahhu/NeuralTraderX/api"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	api.RegisterApiRoutes(router)

	router.Run(":3000")
}
