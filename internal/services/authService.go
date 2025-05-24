package services

import (
	"errors"
	"net/http"
	"rest-api/internal/models"
	"rest-api/internal/repositories"
	"rest-api/internal/types"
	"rest-api/internal/utils"

	"gorm.io/gorm"
)


type AuthService interface {
	Login(payload types.LoginPayload) types.AppResponse
	Signup(payload types.SignupPayload) types.AppResponse
}


type authService struct {
	repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) AuthService {
	return &authService{repo: repo}
}

func (r authService) Login(payload types.LoginPayload) types.AppResponse {
	user, err := r.repo.FindOne(models.User{Email: payload.Email});

	if err != nil {
		return types.AppResponse{Code: http.StatusInternalServerError, Response: types.Response{Message: "Failed to login"}}
	}

	err = utils.CompareHash(payload.Password, user.Password)

	if err != nil {
		return types.AppResponse{Code: http.StatusUnauthorized, Response: types.Response{Message: "Invali email or password"}}
	}

	token, err := utils.GenerateToken(user.Email);
	if err != nil {
		return types.AppResponse{Code: http.StatusInternalServerError, Response: types.Response{Message: "Failed to login"}}
	}

	return types.AppResponse{Code: http.StatusOK, Response: types.Response{
		Message: "User login successfully",
		Data: map[string]interface{}{
			"token": token,
		},
	} }
}


func (r authService) Signup(payload types.SignupPayload) types.AppResponse {
	existingUser, err := r.repo.FindOne(models.User{Email: payload.Email});
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return types.AppResponse{Code: http.StatusInternalServerError, Response: types.Response{Message: "Internal server error"}}
	}

	if err == nil && existingUser != nil {
		return types.AppResponse{Code: http.StatusConflict, Response: types.Response{Message: "User already exists"}}
	}
	
	hashedPassword, err := utils.HashText(payload.Password);
	if err != nil {
		return types.AppResponse{Code: http.StatusInternalServerError, Response: types.Response{Message: "Internal server error"}}
	}

	newUser := models.User{Email: payload.Email, Password: hashedPassword, Username: ""}
	err = r.repo.Create(newUser);
	
	if err != nil {
		return types.AppResponse{Code: http.StatusInternalServerError, Response: types.Response{Message: "Failed to create user"}}
	}

	return types.AppResponse{Code: http.StatusCreated, Response: types.Response{Message: "User created successfully"}}
}