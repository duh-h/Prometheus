package main

import (
	"log"
	"os"

	"github.com/duh-h/Prometheus/api/database"
	"github.com/duh-h/Prometheus/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitializeClient()

	router := gin.Default()

	setupRouters(router)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(router.Run(":" + port))

}

func setupRouters(router *gin.Engine) {

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello",
		})
	})
	router.POST("/api/v1", routes.ShortenURL)
	router.GET("/api/v1/:shortID", routes.GetByShortID)
	router.DELETE("api/v1/:shortID", routes.DeleteURL)
	router.PUT("api/v1/:shortID", routes.EditURL)
	router.POST("api/v1/addTag", routes.AddTag)
}
