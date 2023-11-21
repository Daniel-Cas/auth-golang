package controller

import (
	"auth-golang/domain"
	"auth-golang/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

type UserController struct {
	Service service.UserService
}

/*
func (controller *UserController) GetAll(ctx *gin.Context) {
	users, err := controller.Service.FindAll()

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Error when getting users",
		})
		return
	}
	ctx.JSON(200, users)
}

func (controller *UserController) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	user, err := controller.Service.FindById(id)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Error when getting user",
		})
		return
	}
	ctx.JSON(200, user)
}

*/

func (controller *UserController) Save(ctx *gin.Context) {
	id := uuid.New().String()
	var user domain.User
	err := ctx.BindJSON(&user)
	user.Id = id
	user.CreatedAt = time.Now().GoString()

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Error when binding user",
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
