package controller

import (
	"auth-golang/domain"
	"auth-golang/service"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"time"
)

type UserController struct {
	Service service.UserService
}

func (controller *UserController) GetAll(ctx *gin.Context) {
	users, err := controller.Service.FindAll()

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Error when getting users",
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(200, users)
}

func (controller *UserController) GetById(ctx *gin.Context) {
	idParameter := ctx.Param("id")
	id, err := gocql.ParseUUID(idParameter)
	user, err := controller.Service.FindById(id)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Error when getting user" + err.Error(),
		})
		return
	}
	ctx.JSON(200, user)
}

func SetDependencies(service service.UserService) *UserController {
	return &UserController{Service: service}
}

func (controller *UserController) Save(ctx *gin.Context) {
	var user domain.User
	user.Id = uuid.New().String()
	user.CreatedAt = time.Now()
	err := ctx.BindJSON(&user)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Error when binding user: " + user.Id,
			"error":   err.Error(),
		})
		return
	}

	err = controller.Service.Save(&user)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Error when saving user",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(200, user)
}
