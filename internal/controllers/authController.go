package controllers

import (
	"errors"
	"net/http"
	"rest-api/internal/models"
	"rest-api/internal/services"
	"rest-api/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		return;
	}

	existingUser, err := ac.userService.FindOne(models.User{Email: user.Email});
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return;
	}

	if err == nil && existingUser != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "User already exists",
		})
		return;
	}
	
	hashedPassword, err := utils.HashText(user.Password);
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
		return;
	}

	newUser := models.User{Email: user.Email, Password: hashedPassword, Username: user.Name}
	err = ac.userService.Create(newUser);
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user",
		})
		return;
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
	return;
}

type LoginPayload struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (ac *authController) Login(ctx *gin.Context) {
	var reqPayload LoginPayload
	err := ctx.ShouldBindJSON(&reqPayload);

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Please provide valid request",
		})
		return;
	}

	user, err := ac.userService.FindOne(models.User{Email: reqPayload.Email});

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to login",
		})
		return;
	}

	err = utils.CompareHash(reqPayload.Password, user.Password)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email or password",
		})
		return;
	}

	token, err := utils.GenerateToken(user.Email);
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to login",
		})
		return;
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User login successfully",
		"data": gin.H{ "token": token },
	})
	return;
}
