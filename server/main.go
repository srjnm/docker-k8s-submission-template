package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/srjnm/docker-k8s-submission-template/server/controllers"
	"github.com/srjnm/docker-k8s-submission-template/server/handlers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("API-Server running on port " + port)

	server := gin.Default()

	apiController := controllers.NewApiController()
	apiHandler := handlers.NewApiHandler(apiController)

	server.GET("/", apiHandler.RootHandler)
	server.GET("/cc", apiHandler.CCGetHandler)
	server.POST("/cc", apiHandler.CCPostHandler)

	server.Run(":" + port)
}
