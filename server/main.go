package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/srjnm/docker-k8s-submission-template/server/controllers"
	"github.com/srjnm/docker-k8s-submission-template/server/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := gin.Default()

	apiController := controllers.NewApiController()
	apiHandler := handlers.NewApiHandler(apiController)

	server.GET("/", apiHandler.RootHandler)
	server.GET("/cc", apiHandler.CCGetHandler)

	server.Run(":" + port)
}
