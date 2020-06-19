package api

import (
	"net/http"

	cd "github.com/devrodriguez/multitienda-api/internal/category/domain"
	cud "github.com/devrodriguez/multitienda-api/internal/customer/domain"
	sd "github.com/devrodriguez/multitienda-api/internal/store/domain"
	"github.com/gin-gonic/gin"
)

type ICategoryHandler interface {
	GetCategories(c *gin.Context)
}

type IStoreHandler interface {
	GetStores(c *gin.Context)
}

type ICustomerHandler interface {
	GetAll(c *gin.Context)
	Create(c *gin.Context)
}

type categoryHandler struct {
	serviceContract cd.ServiceContract
}

type storeHandler struct {
	serviceContract sd.ServiceContract
}

type customerHandler struct {
	adapter cud.CustomerPortOut
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

func NewCustomerHandler(adapter cud.CustomerPortOut) ICustomerHandler {
	return &customerHandler{
		adapter,
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

func (h *customerHandler) GetAll(c *gin.Context) {
	customers, err := h.adapter.GetAll()
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, customers)
}

func (h *customerHandler) Create(c *gin.Context) {
	var customer cud.Customer
	c.BindJSON(&customer)
	err := h.adapter.Create(customer)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}
