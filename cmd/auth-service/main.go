package main

import (
<<<<<<< HEAD
<<<<<<< Updated upstream
	"log"
	"net/http"
=======
	cfg "auth-service/internal/config"
	logger "auth-service/internal/middleware"
	"auth-service/internal/routes"
>>>>>>> Stashed changes
=======
	"auth-service/internal/config"
	logger "auth-service/internal/middleware"
	"auth-service/internal/routes"
>>>>>>> b2f126e71dd529932cc2e399b8239a71d17932d1

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

<<<<<<< HEAD
<<<<<<< Updated upstream
type query struct{}

func (_ *query) Hello() string {
	return "Hello, world!"
}
=======
func main() {
	cfg.ConfigInit()
	log.SetLevel(log.DebugLevel)
	cfg.ConnectDB()
	router := gin.Default()
>>>>>>> Stashed changes

=======
>>>>>>> b2f126e71dd529932cc2e399b8239a71d17932d1
func main() {
	log.SetLevel(log.DebugLevel)
	config.ConnectDB()
	router := gin.Default()

	router.Use(logger.GinLogger())
	routes.RegisterRoutes(router)

	logger.Log.Info("GraphQL server starting on http://localhost:8085/graphql")

	router.Run(":8080")
}
