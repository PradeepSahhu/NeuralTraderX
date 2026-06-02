package api

import (
	"github.com/gin-gonic/gin"
)

type RouteRegistrar interface {
	RegisterRoutes(*gin.RouterGroup)
}
