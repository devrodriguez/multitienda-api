package main

import (
	"io"
	"log"
	"os"

	"github.com/devrodriguez/multitienda-api/controllers"
	"github.com/devrodriguez/multitienda-api/db"
	"github.com/devrodriguez/multitienda-api/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set environment variables
	os.Setenv("JWT_SECRET", "dev1986")
	os.Setenv("EXPIRATION", "30m")

	port := os.Getenv("PORT")
	server := gin.New()

	// Set log file
	setupLogOutput()

	// Middlewares
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.CORSAllowed())

	// Db connection
	connErr := db.Connect()
	if connErr != nil {
		log.Fatal(connErr)
		return
	}

	// Routes
	pubRoutes := server.Group("/api")
	{
		pubRoutes.GET("/signin", controllers.SignIn)
		pubRoutes.GET("/stores", controllers.GetStores)
		pubRoutes.POST("/stores", controllers.CreateStore)
	}

	if port == "" {
		port = "3001"
	}

	// Run server
	server.Run(":" + port)
}

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
