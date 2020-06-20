package main

import (
	"log"

	api "github.com/devrodriguez/multitienda-api/cmd/api"
	ca "github.com/devrodriguez/multitienda-api/internal/category/application"
	ci "github.com/devrodriguez/multitienda-api/internal/category/infrastructure"
	customer "github.com/devrodriguez/multitienda-api/internal/customer"
	sa "github.com/devrodriguez/multitienda-api/internal/store/application"
	si "github.com/devrodriguez/multitienda-api/internal/store/infrastructure"
	"github.com/devrodriguez/multitienda-api/middlewares"
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

	cusPortOut, err := customer.NewAdapterOut(mongoURL, mongoDB, mongoTimeout)
	if err != nil {
		log.Fatal(err)
	}

	cs := ca.NewCategoryService(cr)
	ss := sa.NewStoreService(sr)
	ch := api.NewCategoryHandler(cs)
	sh := api.NewStoreHandler(ss)
	cusAdapIn := customer.NewAdapterIn(cusPortOut)
	cuh := api.NewCustomerHandler(cusAdapIn)

	app := gin.New()
	app.Use(gin.Recovery(), middlewares.Logger(), middlewares.CORSAllowed())

	app.GET("/categories", ch.GetCategories)
	app.GET("/stores", sh.GetStores)
	app.GET("/customers", cuh.GetAll)
	app.POST("/customers", cuh.Create)
	app.Run(":3001")
}
