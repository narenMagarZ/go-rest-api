package controllers

import (
	"net/http"
	"rest-api/internal/models"
	"rest-api/internal/services"
	"rest-api/internal/utils"

	"github.com/gin-gonic/gin"
)


type authController struct {
	userService services.UserService
}

func NewAuthController(userSerivce services.UserService) *authController {
	return &authController{userService: userSerivce}
}

func (ac *authController) Signup(ctx *gin.Context) {
	type User struct {
		Name string `json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
	}
	var user User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
	}

	existingUser, err := ac.userService.FindOne(models.User{Email: user.Email});
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
	}

	if existingUser != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "User already exists",
		})
	}
	
	hashedPassword, err := utils.HashText(user.Password);
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
	}

	newUser := models.User{Email: user.Email, Password: hashedPassword, Username: ""}
	err = ac.userService.Create(newUser);
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user",
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}

func (ac *authController) Login(ctx *gin.Context) {
	// accept user email/username and password
	// check if database
	// verify accessToken
	// verify refereshToken in case of accessToken throws expire token
}
