package routes

import (
	"auth-service/internal/graphql"
	glq "auth-service/internal/graphql"
	mdl "auth-service/internal/middleware"
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
		// This step ensures that the context includes what's been set in Gin's context
		ctx := context.WithValue(c.Request.Context(), "resolver", c.MustGet("resolver"))
		rWithCtx := c.Request.WithContext(ctx)
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
	r.Use(mdl.AuthMiddleware())
	r.POST("/graphql", GraphQLHandler())
	// router.POST("/graphql", gin.WrapH(h))
	r.GET("/graphql", gin.WrapH(h))
}
