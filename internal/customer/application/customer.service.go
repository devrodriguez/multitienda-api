package customer

import (
	domain "github.com/devrodriguez/multitienda-api/internal/customer/domain"
)

type adapterOut struct {
	adapterIn domain.CustomerPortIn
}

func NewCustomerService(adapterIn domain.CustomerPortIn) domain.CustomerPortOut {
	return &adapterOut{
		adapterIn: adapterIn,
	}
}

func (ao *adapterOut) GetAll() ([]*domain.Customer, error) {
	return ao.adapterIn.GetAll()
}

func (ao *adapterOut) Create(customer domain.Customer) error {
	return ao.adapterIn.Create(customer)
}
