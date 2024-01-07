package route

import (
	"auth-golang/config"
	"auth-golang/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, services *config.Services) {
	userController := controller.SetDependencies(*services.User)

	router.GET("/users", userController.GetAll)
	router.GET("/user/:id", userController.GetById)
	router.POST("/user", userController.Save)
}
