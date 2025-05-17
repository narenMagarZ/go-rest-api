package controllers

import (
	"net/http"
	"rest-api/internal/services"
	"strconv"

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
	id, err := strconv.Atoi(ctx.Param("id"));
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id",
		})
	}

	user, err := c.userService.FindById(id);
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
	}

	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User fetched successfully",
		"data": user,
	})
}

func (c *userController) UpdateUser(ctx *gin.Context) {}

func (c *userController) GetAllUsers(ctx *gin.Context) {

}

func (c *userController) DeleteUser(ctc *gin.Context) {

}
