package main

import (
	"chat_me_api/src/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	app := gin.New()
	app.Use(gin.Logger())
	app.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))

	// Pipeline
	routes.AuthRoutes(app)

	app.Run(":" + port)
}
