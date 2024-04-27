package middleware

import (
	cfg "auth-service/internal/config"
	"auth-service/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Extract the Authorization header.
// 		tokenString := c.GetHeader("Authorization")

// 		// If no Authorization header is present, just call the next handler.
// 		if tokenString == "" {
// 			c.Next()
// 			return
// 		}

// 		// Continue with token validation if the header is present.
// 		splitToken := strings.Split(tokenString, "Bearer ")
// 		if len(splitToken) != 2 {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization token format"})
// 			return
// 		}
// 		tokenString = splitToken[1]

// 		_, err := utils.ParseToken(tokenString, cfg.SaltPassKey)
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
// 			return
// 		}

// 		// Set the user ID in context and continue.
// 		// c.Set("userID", userID)
// 		c.Next()
// 	}
// }

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userAuthorized := false
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.Set("authorized", userAuthorized)
			c.Next()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		userID, err := utils.ParseToken(tokenString, cfg.SaltPassKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"isLoggedIn": false, "message": "Error on token ", "userId": userID})
			return
		}
		userAuthorized = true
		c.Set("userID", userID)
		c.Set("authorized", userAuthorized)
		c.Next()
		// c.AbortWithStatusJSON(http.StatusOK, gin.H{"isLoggedIn": true, "message": "Already authenticated via token", "userId": userID})
	}
}
