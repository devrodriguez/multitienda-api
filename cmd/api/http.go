package api

import (
	"net/http"

	category "github.com/devrodriguez/multitienda-api/category/domain"
	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	GetCategories(c *gin.Context)
}

type handler struct {
	serviceContract category.ServiceContract
}

func NewHandler(serviceContract category.ServiceContract) CategoryHandler {
	return &handler{
		serviceContract: serviceContract,
	}
}

func (h *handler) GetCategories(c *gin.Context) {

	categories, err := h.serviceContract.GetAll()

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, categories)
}
