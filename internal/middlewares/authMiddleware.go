package middlewares

import (
	"fmt"
	"net/http"
	"rest-api/internal/models"
	"rest-api/internal/services"
	"rest-api/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func (c *gin.Context)  {
		c.Next()
	}
}

func Authenticate(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		token := strings.Replace(bearerToken, "Bearer ", "", 1)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized user",
			})
			return
		}
		claims, err := utils.VerifyToken(token)
		if err != nil {
			fmt.Println("Failed to verify token")
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized user"})
			return
		}
		userExists, err := userService.FindOne(models.User{Email: claims.Email})
		if err != nil {
			fmt.Println("User doesn't exist")
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized user"})
		}
		c.Set("user", userExists);
		c.Next();
	}
}


