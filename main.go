package main

import (
	"golang-rest-api/config"
	"golang-rest-api/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDb()
}

func main() {
	router := gin.New()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Server is live..."})
	})

	routes.UserRoute(router)

	router.Run()
}
