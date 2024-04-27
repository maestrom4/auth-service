package middleware

import (
	"net/http"
	"strings"

	cfg "auth-service/internal/config"
	"auth-service/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.Next()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		userID, err := utils.ParseToken(tokenString, cfg.SaltPassKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"isLoggedIn": false, "message": "Error on token bearer", "userId": userID})
			return
		}

		c.Set("userID", userID)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{"isLoggedIn": true, "message": "Already authenticated via token", "userId": userID})
	}
}
