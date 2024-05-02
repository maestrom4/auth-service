package routes

import (
	cfg "auth-service/internal/config"
	"auth-service/pkg/mongodb"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyEmail() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.Query("token")
		userRepository := mongodb.NewUserRepository(cfg.GetDBCollection(cfg.CollectionUser))
		_, err := userRepository.VerifyUserByToken(c, token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"isSuccess": true, "message": "Email verified successfully!"})
	}

}
