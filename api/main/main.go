package main

import (
	"auth-golang/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	userController := controller.UserController{}

	router.GET("/users", userController.GetAll)
	router.GET("/user/:id", userController.GetById)
	router.POST("/user", userController.Save)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}
