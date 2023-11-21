package route

import (
	"auth-golang/controller"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	userController := controller.UserController{}

	//router.GET("/users", userController.GetAll)
	//router.GET("/user/:id", userController.GetById)
	router.POST("/user", userController.Save)
}
