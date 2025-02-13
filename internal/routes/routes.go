package routes

import (
	v "auth-service/internal/config"
	"auth-service/internal/graphql"
	glq "auth-service/internal/graphql"
	mdl "auth-service/internal/middleware"

	u "auth-service/utils"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func RegisterRoutes(r *gin.Engine) {
	h := handler.New(&handler.Config{
		Schema:   &glq.Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	r.Use(mdl.SecureHeadersMiddleware())
	r.Use(mdl.ResolverMiddleware())
	r.Use(mdl.AuthMiddleware(u.ParseToken))
	r.POST("/graphql", GraphQLHandler())
	// router.POST("/graphql", gin.WrapH(h))
	r.GET("/graphql", gin.WrapH(h))
	r.GET("/verify", VerifyEmail())
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
