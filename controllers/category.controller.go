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

func GetCategories(gCtx *gin.Context) {
	var categories []*models.Category
	var response models.Response

	mgClient := *db.GetClient()
	findOptions := options.Find()

	//findOptions.SetLimit(2)

	categoriesRef := mgClient.Database("multitienda").Collection("categories")
	categoriesCur, err := categoriesRef.Find(context.TODO(), bson.D{{}}, findOptions)

	if err != nil {
		response.Message = "Error getting data"
		response.Error = err.Error()

		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	for categoriesCur.Next(context.TODO()) {
		var category models.Category

		err := categoriesCur.Decode(&category)
		if err != nil {
			log.Fatal(err)
		}

		categories = append(categories, &category)
	}

	if err := categoriesCur.Err(); err != nil {
		log.Fatal(err)
	}

	categoriesCur.Close(context.TODO())

	gCtx.JSON(http.StatusOK, categories)
}
