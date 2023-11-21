package route

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	routes := router.Group("/api/auth")

	UserRoutes(routes)

	routes.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"response": "ok",
		})
	})

	return router
}
