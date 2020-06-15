package main

import (
	"log"

	application "github.com/devrodriguez/multitienda-api/category/application"
	infrastructure "github.com/devrodriguez/multitienda-api/category/infrastructure"
	api "github.com/devrodriguez/multitienda-api/cmd/api"
	"github.com/devrodriguez/multitienda-api/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	mongoURL := "mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/test?retryWrites=true&w=majority"
	mongoDB := "multitienda"
	mongoTimeout := 10

	repo, err := infrastructure.NewMongoRepository(mongoURL, mongoDB, mongoTimeout)
	if err != nil {
		log.Fatal(err)
	}

	service := application.NewCategoryService(repo)
	handler := api.NewHandler(service)

	app := gin.New()
	app.Use(gin.Recovery(), middlewares.Logger(), middlewares.CORSAllowed())

	app.GET("/categories", handler.GetCategories)
	app.Run(":3001")
}
