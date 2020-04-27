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

func GetCustomers(gCtx *gin.Context) {
	var customers []*models.Customer
	var response models.Response

	mgClient := *db.GetClient()
	findOptions := options.Find()

	customersRef := mgClient.Database("multitienda").Collection("customers")
	customersCur, err := customersRef.Find(context.TODO(), bson.D{{}}, findOptions)

	if err != nil {
		response.Message = "Error getting data"
		response.Error = err.Error()

		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	for customersCur.Next(context.TODO()) {
		var customer models.Customer

		err := customersCur.Decode(&customer)
		if err != nil {
			log.Fatal(err)
		}

		customers = append(customers, &customer)
	}

	if err := customersCur.Err(); err != nil {
		log.Fatal(err)
	}

	customersCur.Close(context.TODO())

	gCtx.JSON(http.StatusOK, customers)
}

func CreateCustomer(gCtx *gin.Context) {
	var customer models.Customer
	var response models.Response

	mgClient := *db.GetClient()

	// Binding data request
	if err := gCtx.BindJSON(&customer); err != nil {
		response.Error = err.Error()
		response.Message = "Error binding data"
		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	customerRef := mgClient.Database("multitienda").Collection("customers")
	insRes, err := customerRef.InsertOne(context.TODO(), customer)

	if err != nil {
		response.Error = err.Error()
		response.Message = "Error setting data"
		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	// Build response
	response.Data = gin.H{"docID": insRes.InsertedID}
	response.Message = "Document created"

	gCtx.JSON(http.StatusOK, response)

}
