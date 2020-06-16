package main

import (
	"log"

	ca "github.com/devrodriguez/multitienda-api/category/application"
	ci "github.com/devrodriguez/multitienda-api/category/infrastructure"
	api "github.com/devrodriguez/multitienda-api/cmd/api"
	"github.com/devrodriguez/multitienda-api/middlewares"
	sa "github.com/devrodriguez/multitienda-api/store/application"
	si "github.com/devrodriguez/multitienda-api/store/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	mongoURL := "mongodb+srv://adminUser:Chrome.2020@auditcluster-ohkrf.gcp.mongodb.net/test?retryWrites=true&w=majority"
	mongoDB := "multitienda"
	mongoTimeout := 10

	cr, err := ci.NewMongoRepository(mongoURL, mongoDB, mongoTimeout)
	if err != nil {
		log.Fatal(err)
	}

	sr, err := si.NewMongoRepository(mongoURL, mongoDB, mongoTimeout)
	if err != nil {
		log.Fatal(err)
	}

	cs := ca.NewCategoryService(cr)
	ss := sa.NewStoreService(sr)
	ch := api.NewCategoryHandler(cs)
	sh := api.NewStoreHandler(ss)

	app := gin.New()
	app.Use(gin.Recovery(), middlewares.Logger(), middlewares.CORSAllowed())

	app.GET("/categories", ch.GetCategories)
	app.GET("/stores", sh.GetStores)
	app.Run(":3001")
}
