package main

import (
	"latihan5-jwt/auth"
	"latihan5-jwt/config"
	"latihan5-jwt/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDatabase()
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		var foundUser models.User
		result := db.Where("username = ? AND password = ?", user.Username, user.Password).First(&foundUser)
		if result.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
			return
		}

		token, err := auth.GenerateToken(user.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while generating token"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	r.GET("/protected", AuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Protected content"})
	})

	r.Run(":8080")
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		token, err := auth.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
			c.Abort()
			return
		}
		c.Next()
	}
}
