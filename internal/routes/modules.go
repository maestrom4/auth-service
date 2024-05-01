package routes

import (
	cfg "auth-service/internal/config"
	v "auth-service/internal/config"
	"auth-service/internal/graphql"
	"auth-service/pkg/mongodb"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
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

func GraphQLHandler() gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:   &graphql.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return func(c *gin.Context) {
		resolver := c.MustGet(string(v.ResolverKey))
		userID := c.MustGet(string(v.UserIDKey))
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, string(v.ResolverKey), resolver)
		ctx = context.WithValue(ctx, string(v.UserIDKey), userID)
		rWithCtx := c.Request.WithContext(ctx)
		h.ServeHTTP(c.Writer, rWithCtx)
	}
}
