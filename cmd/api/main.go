package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

	"rest-api/internal/config"
	"rest-api/internal/controllers"
	"rest-api/internal/middlewares"
	"rest-api/internal/models"
	"rest-api/internal/repositories"
	"rest-api/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	router := gin.Default()
	router.Run(fmt.Sprintf(":%s", port))

	db := config.ConnectDB()

	db.AutoMigrate(models.User{});

	userRepository := repositories.NewUserRepository(db);
	userService := services.NewUserService(userRepository);

	userController := controllers.NewUserController(userService);

	api := router.Group("/api/v1")
	api.Use(middlewares.Logger());
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login")
			auth.POST("/register")
		}

		users := api.Group("/users")
		users.Use(middlewares.Authenticate());
		{
			users.GET("/", userController.GetAllUsers)
			users.POST("/", userController.CreateUser)
			users.GET("/:id", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "User fetched successfully",
					"status": true,
				})
			})
			users.PUT("/:id", userController.UpdateUser)
			users.DELETE("/:id", userController.DeleteUser)
		}

	}
}