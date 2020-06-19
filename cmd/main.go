package main

import (
	"log"

	api "github.com/devrodriguez/multitienda-api/cmd/api"
	ca "github.com/devrodriguez/multitienda-api/internal/category/application"
	ci "github.com/devrodriguez/multitienda-api/internal/category/infrastructure"
	cua "github.com/devrodriguez/multitienda-api/internal/customer/application"
	cui "github.com/devrodriguez/multitienda-api/internal/customer/infrastructure"
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

	portIn, err := cui.NewMongoAdapter(mongoURL, mongoDB, mongoTimeout)

	cs := ca.NewCategoryService(cr)
	ss := sa.NewStoreService(sr)
	ch := api.NewCategoryHandler(cs)
	sh := api.NewStoreHandler(ss)
	cus := cua.NewCustomerService(portIn)
	cuh := api.NewCustomerHandler(cus)

	app := gin.New()
	app.Use(gin.Recovery(), middlewares.Logger(), middlewares.CORSAllowed())

	app.GET("/categories", ch.GetCategories)
	app.GET("/stores", sh.GetStores)
	app.GET("/customers", cuh.GetAll)
	app.POST("/customers", cuh.Create)
	app.Run(":3001")
}
