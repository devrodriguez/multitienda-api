package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/devrodriguez/multitienda-api/db"
	"github.com/devrodriguez/multitienda-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetStores(gCtx *gin.Context) {
	var stores []*models.Store
	var response models.Response

	mgClient := *db.GetClient()
	findOptions := options.Find()

	//findOptions.SetLimit(2)

	storesRef := mgClient.Database("multitienda").Collection("stores")
	storesCur, err := storesRef.Find(context.TODO(), bson.D{{}}, findOptions)

	if err != nil {
		response.Message = "Error getting data"
		response.Error = err.Error()

		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	for storesCur.Next(context.TODO()) {
		var store models.Store

		err := storesCur.Decode(&store)
		if err != nil {
			log.Fatal(err)
		}

		stores = append(stores, &store)
	}

	if err := storesCur.Err(); err != nil {
		log.Fatal(err)
	}

	storesCur.Close(context.TODO())

	gCtx.JSON(http.StatusOK, stores)
}

func CreateStore(gCtx *gin.Context) {
	var response models.Response
	var store models.Store
	mgClient := db.GetClient()

	if err := gCtx.BindJSON(&store); err != nil {
		response.Error = err.Error()
		response.Message = "Error binding JSON data"
		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	storeRef := mgClient.Database("multitienda").Collection("stores")
	insertRes, err := storeRef.InsertOne(context.TODO(), store)

	if err != nil {
		response.Error = err.Error()
		response.Message = "Error setting data"
		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	// Build response
	response.Message = "Document created"
	response.Data = gin.H{"docID": insertRes.InsertedID}
	gCtx.JSON(http.StatusOK, response)
}
