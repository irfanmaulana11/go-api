package server

import (
	"go-api/helpers"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// Middleware function for origin permit
func CORSMiddleware(c *gin.Context) {
	origins := os.Getenv("ALLOWED_ORIGINS")

	c.Header("Access-Control-Allow-Origin", origins)
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

// Middleware function for token validation
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")

		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Missing authorization header"})
			c.Abort()
			return
		}

		// Check if the Authorization header starts with "Bearer "
		if !strings.HasPrefix(authorizationHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization format"})
			c.Abort()
			return
		}

		// Extract the token part from the Authorization header
		tokenString := authorizationHeader[len("Bearer "):]

		if err := helpers.ValidateToken(tokenString); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}
