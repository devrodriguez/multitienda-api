package api

import (
	"net/http"

	cd "github.com/devrodriguez/multitienda-api/category/domain"
	sd "github.com/devrodriguez/multitienda-api/store/domain"
	"github.com/gin-gonic/gin"
)

type ICategoryHandler interface {
	GetCategories(c *gin.Context)
}

type IStoreHandler interface {
	GetStores(c *gin.Context)
}

type categoryHandler struct {
	serviceContract cd.ServiceContract
}

type storeHandler struct {
	serviceContract sd.ServiceContract
}

func NewCategoryHandler(serviceContract cd.ServiceContract) ICategoryHandler {
	return &categoryHandler{
		serviceContract: serviceContract,
	}
}

func NewStoreHandler(serviceContract sd.ServiceContract) IStoreHandler {
	return &storeHandler{
		serviceContract: serviceContract,
	}
}

func (h *categoryHandler) GetCategories(c *gin.Context) {

	categories, err := h.serviceContract.GetAll()

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, categories)
}

func (h *storeHandler) GetStores(c *gin.Context) {
	stores, err := h.serviceContract.GetAllStores()

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, stores)
}
