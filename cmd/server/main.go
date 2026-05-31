package main

import (
	"github.com/PradeepSahhu/NeuralTraderX/api"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	api.RegisterHealthApiRoutes(router)
	api.RegisterBasicRoutes(router)

	router.Run(":3000")
}
