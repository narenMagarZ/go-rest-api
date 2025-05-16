package controllers

import (
	"rest-api/internal/services"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type userController struct {
	userService services.UserService
}

func NewUserController(service services.UserService) UserController {
	return &userController{userService: service}
}

func (c *userController) CreateUser(ctx *gin.Context) {

}

func (c *userController) GetUser(ctx *gin.Context) {
}

func (c *userController) UpdateUser(ctx *gin.Context) {}

func (c *userController) GetAllUsers(ctx *gin.Context) {

}

func (c *userController) DeleteUser(ctc *gin.Context) {

}
