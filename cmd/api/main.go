package main

import (
	"fmt"
	"log"
	"os"

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

	db := config.ConnectDB()

	db.AutoMigrate(models.User{});

	userRepository := repositories.NewUserRepository(db);
	userService := services.NewUserService(userRepository);

	api := router.Group("/api/v1")
	api.Use(middlewares.Logger());
	{
		auth := api.Group("/auth")
		authController := controllers.NewAuthController(userService);
		{
			auth.POST("/login", authController.Login)
			auth.POST("/signup", authController.Signup)
		}

		users := api.Group("/users")
		users.Use(middlewares.Authenticate());
		{
			
		
			userController := controllers.NewUserController(userService);

			users.GET("/", userController.GetAllUsers)
			users.POST("/", userController.CreateUser)
			users.GET("/:id", userController.GetUser)
			users.PUT("/:id", userController.UpdateUser)
			users.DELETE("/:id", userController.DeleteUser)
		}

	}

	router.Run(fmt.Sprintf(":%s", port))
}