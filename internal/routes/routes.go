package routes

import (
	"auth-service/internal/graphql"
	glq "auth-service/internal/graphql"
	mdlwr "auth-service/internal/middleware"
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

func RegisterRoutes(router *gin.Engine) {
	h := handler.New(&handler.Config{
		Schema:   &glq.Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	router.Use(mdlwr.ResolverMiddleware())
	router.POST("/graphql", GraphQLHandler())
	// router.POST("/graphql", gin.WrapH(h))
	router.GET("/graphql", gin.WrapH(h))
}
