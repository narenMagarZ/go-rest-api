package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"rest-api/internal/models"
	"rest-api/internal/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userController struct {
	userService services.UserService
}

func NewUserController(service services.UserService) *userController {
	return &userController{userService: service}
}

func (c *userController) GetUser(ctx *gin.Context) {
	paramId := ctx.Param("id")
	if paramId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Please provide user id."})
		return;
	}
	id, err := strconv.Atoi(paramId);
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user id",
		})
		return;
	}

	user, err := c.userService.FindById(id);
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
			return;
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return;
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User fetched successfully",
		"data": user,
	})
	return;
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	paramId := ctx.Param("id")
	if paramId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Please provide user id."})
		return;
	}

	id, err := strconv.Atoi(paramId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user id."})
		return;
	}
	_, err = c.userService.FindById(id);
	
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"});
			return;
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update user"})
		return;
	}

	data, err := io.ReadAll(ctx.Request.Body)
	var payload models.User
	err = json.Unmarshal(data, &payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update user"})
		return;
	}
	err = c.userService.UpdateOne(id, payload)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update user"})
		return;
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"});
	return;
}

func (c *userController) GetAllUsers(ctx *gin.Context) {
}

func (c *userController) DeleteUser(ctx *gin.Context) {
	paramId := ctx.Param("id")
	if paramId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Please provide user id"})
		return;
	}
	id, err := strconv.Atoi(paramId);
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid user id"})
		return;
	}
	_, err = c.userService.FindById(id);
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
			return;
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete user"})
	}
	err = c.userService.DeleteOne(struct{Id int}{Id: id});
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete user",
		})
		return;
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"});
	return;
}
