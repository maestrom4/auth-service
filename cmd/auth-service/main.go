package main

import (
	cfg "auth-service/internal/config"
	logger "auth-service/internal/middleware"
	"auth-service/internal/routes"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	cfg.ConfigInit()
}
func main() {
	// cfg.ConfigInit()
	cfg.InitLogger()
	log.SetLevel(log.DebugLevel)
	cfg.ConnectDB()
	router := gin.Default()

	router.Use(logger.GinLogger())
	routes.RegisterRoutes(router)

	logger.Log.Info("GraphQL server starting on http://localhost:8085/graphql")

	router.Run(":8080")
}
