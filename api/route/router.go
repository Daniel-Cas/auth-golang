package route

import (
	"auth-golang/config"
	"github.com/gin-gonic/gin"
)

func SetupRouter(services *config.Services) *gin.Engine {
	router := gin.Default()

	routes := router.Group("/api/auth")

	UserRoutes(routes, services)

	routes.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"response": "ok",
		})
	})

	return router
}
