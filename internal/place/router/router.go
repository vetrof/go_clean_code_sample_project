package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	nearbyHandler gin.HandlerFunc,
	popularHandler gin.HandlerFunc,
) *gin.Engine {
	r := gin.Default()
	r.GET("/nearby", nearbyHandler)
	r.GET("/popular", popularHandler)
	return r
}
