package middleware

import (
	cfg "auth-service/internal/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(parseTokenFunc func(string, string) (string, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(string(cfg.UserIDKey), "")
		tokenString := c.GetHeader("Authorization")
		if strings.HasPrefix(tokenString, "Bearer ") {
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
			userID, err := parseTokenFunc(tokenString, cfg.JwtSecretKey)
			if err == nil {
				c.Set(string(cfg.UserIDKey), userID)
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
				return
			}
		}
		c.Next()
	}
}
