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

func GraphQLHandler() gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:   &graphql.Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return func(c *gin.Context) {
		// userID := c.MustGet(string(v.UserIDKey))
		// ctx := context.WithValue(c.Request.Context(), string(v.ResolverKey), c.MustGet(string(v.ResolverKey)))
		// ctx = context.WithValue(c.Request.Context(), v.UserIDKey, userID)
		// rWithCtx := c.Request.WithContext(ctx)
		// h.ServeHTTP(c.Writer, rWithCtx)

		// ctx := context.WithValue(c.Request.Context(), string(v.ResolverKey), c.MustGet(string(v.ResolverKey)))
		// rWithCtx := c.Request.WithContext(ctx)
		// h.ServeHTTP(c.Writer, rWithCtx)
		resolver := c.MustGet(string(v.ResolverKey))

		// Retrieve the userID from Gin's context
		userID := c.MustGet(string(v.UserIDKey))

		// Start with the request's current context
		ctx := c.Request.Context()

		// Add the resolver to the context
		ctx = context.WithValue(ctx, string(v.ResolverKey), resolver)

		// Add the userID to the context
		ctx = context.WithValue(ctx, string(v.UserIDKey), userID)

		// Create a new request with the updated context
		rWithCtx := c.Request.WithContext(ctx)

		// Serve the HTTP request with the new context
		h.ServeHTTP(c.Writer, rWithCtx)
	}
}

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
}
