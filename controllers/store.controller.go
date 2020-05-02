package controllers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/devrodriguez/multitienda-api/db"
	"github.com/devrodriguez/multitienda-api/models"
	"github.com/devrodriguez/multitienda-api/utilities"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetStores(gCtx *gin.Context) {
	var stores []*models.Store
	var response models.Response

	mgClient := *db.GetClient()
	findOptions := options.Find()

	//findOptions.SetLimit(2)

	storesRef := mgClient.Database("multitienda").Collection("stores")
	storesCur, err := storesRef.Find(context.TODO(), bson.M{"status": "approved"}, findOptions)

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

func FindStores(gCtx *gin.Context) {
	var stores []*models.Store
	var response models.Response
	var posOrigin models.GpsPoint

	mgClient := *db.GetClient()
	findOptions := options.Find()
	q := gCtx.Query("q")

	// Find approved and name ocurence
	filter := bson.M{"status": "approved", "name": primitive.Regex{Pattern: q, Options: ""}}

	log.Println("Lat: ", gCtx.Query("lat"), "Lon: ", gCtx.Query("lon"), "dist: ", gCtx.Query("dist"))

	lat, err := strconv.ParseFloat(gCtx.Query("lat"), 64)
	if err != nil {
		gCtx.JSON(http.StatusOK, stores)
		return
	}

	lon, err := strconv.ParseFloat(gCtx.Query("lon"), 64)
	if err != nil {
		gCtx.JSON(http.StatusOK, stores)
		return
	}

	posOrigin.Lat = lat
	posOrigin.Lon = lon

	//findOptions.SetLimit(2)
	kmDistFixed, err := strconv.ParseFloat(gCtx.Query("dist"), 64)
	if err != nil {
		kmDistFixed = 0
	}

	log.Println("Lat: ", lat, "Lon: ", lon, "dist: ", kmDistFixed)

	storesRef := mgClient.Database("multitienda").Collection("stores")
	storesCur, err := storesRef.Find(context.TODO(), filter, findOptions)

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

		_, kmDist := utilities.HaversineFormule(posOrigin, store.GeoLocation)

		log.Println("Distancia entre puntos: ", kmDist)

		if kmDist <= kmDistFixed {
			stores = append(stores, &store)
		}
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

	store.Status = "pendding"

	// Store images
	pathImgs, err := saveImages(store.String64Images)
	if err != nil {
		response.Message = "Error reading images"
		response.Error = err.Error()
		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	store.String64Images = nil
	store.UrlImages = pathImgs

	// Database reference
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

func saveImages(bs64Img []string) ([]string, error) {
	var pathImgs []string

	for _, v := range bs64Img {
		randStr := utilities.RandomString(12)
		filename := randStr
		filePath, err := utilities.ToFile(v, filename)

		if err != nil {
			log.Println("Error al guardar imagen")
			return nil, err
		}

		pathImgs = append(pathImgs, filePath)
		log.Println(pathImgs)
	}

	return pathImgs, nil
}
