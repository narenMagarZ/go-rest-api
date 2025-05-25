package controllers

import (
	"net/http"
	"rest-api/internal/services"
	"rest-api/internal/types"

	"github.com/gin-gonic/gin"
)

type authController struct {
	userService services.UserService
	authService services.AuthService
}

func NewAuthController(userSerivce services.UserService, authService services.AuthService) *authController {
	return &authController{userService: userSerivce, authService: authService}
}

func (ac *authController) Signup(ctx *gin.Context) {
	var payload types.SignupPayload
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}

	response := ac.authService.Signup(payload)
	ctx.JSON(response.Code, response.Response)
	return
}

func (ac *authController) Login(ctx *gin.Context) {
	var reqPayload types.LoginPayload
	err := ctx.ShouldBindJSON(&reqPayload)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Please provide valid request",
		})
		return
	}

	response := ac.authService.Login(reqPayload)

	ctx.JSON(response.Code, response.Response)
	return
}
