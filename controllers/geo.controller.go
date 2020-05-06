package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/devrodriguez/multitienda-api/models"
	"github.com/gin-gonic/gin"
)

func GetAddresPredictions(gCtx *gin.Context) {
	var response models.Response
	var key = "AIzaSyA9UqT-ykyMUf2MJbc7EBsIQj6D69DHSJo"
	var apiRes models.GoogleApiAuto
	term := url.QueryEscape(gCtx.Query("term"))

	res, err := http.Get("https://maps.googleapis.com/maps/api/place/autocomplete/json?input=" + term + "&key=" + key)

	if err != nil {
		response.Error = err.Error()
		response.Message = "Error requesting data"
		gCtx.JSON(http.StatusInternalServerError, response)
		return
	}

	json.NewDecoder(res.Body).Decode(&apiRes)

	response.Data = apiRes.Predictions
	response.Message = "Success"

	gCtx.JSON(http.StatusOK, response)
}

func AddressToCoordinates(gc *gin.Context) {
	var response models.Response
	address := url.QueryEscape(gc.Query("address"))

	res, err := RequestCoordinates(address)
	if err != nil {
		response.Error = err.Error()
		response.Message = "Error requesting data"
		gc.JSON(http.StatusInternalServerError, response)
		return
	}

	var resModel map[string]interface{}
	json.NewDecoder(res.Body).Decode(&resModel)

	response.Data = resModel
	response.Message = "Success"

	gc.JSON(http.StatusOK, response)
}

func CoordinatesToAddres(gc *gin.Context) {
	var response models.Response
	var resModel map[string]interface{}
	coord := url.QueryEscape(gc.Query("latlng"))

	res, err := RequestAddress(coord)
	if err != nil {
		response.Error = err.Error()
		response.Message = "Error requesting data"
		gc.JSON(http.StatusInternalServerError, response)
		return
	}

	json.NewDecoder(res.Body).Decode(&resModel)

	response.Data = resModel
	response.Message = "Success"

	gc.JSON(http.StatusOK, response)
}

func RequestAddress(coord string) (*http.Response, error) {
	var key = "AIzaSyBnh1nBPOMEml3EbvkzGn0c-LYEKMfKhfE"
	return http.Get("https://maps.googleapis.com/maps/api/geocode/json?latlng=" + coord + "&key=" + key)
}

func RequestCoordinates(address string) (*http.Response, error) {
	var key = "AIzaSyBnh1nBPOMEml3EbvkzGn0c-LYEKMfKhfE"
	return http.Get("https://maps.googleapis.com/maps/api/geocode/json?address=" + address + "&key=" + key)
}
