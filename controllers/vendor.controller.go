package controllers

import (
	"net/http"

	"github.com/devrodriguez/multitienda-api/models"
	"github.com/gin-gonic/gin"
)

func CreateVendor(gCtx *gin.Context) {
	var vendor models.Vendor
	var response models.Response

	// Binding data request
	if err := gCtx.BindJSON(&vendor); err != nil {
		response.Error = err.Error()
		response.Message = "Error binding data"
		gCtx.JSON(http.StatusInternalServerError, response)
	}

}
