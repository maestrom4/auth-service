package logger

import (
	cfg "auth-service/internal/config"
	"auth-service/internal/graphql"
	"auth-service/pkg/mongodb"

	"github.com/gin-gonic/gin"
)

func ResolverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		resolver := &graphql.Resolver{
			UserRepository: mongodb.NewUserRepository(cfg.GetDBCollection(cfg.CollectionUser)),
		}
		c.Set("resolver", resolver)
		c.Next()
	}
}
