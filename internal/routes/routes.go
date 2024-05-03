package routes

import (
	glq "auth-service/internal/graphql"
	mdl "auth-service/internal/middleware"

	u "auth-service/utils"

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
