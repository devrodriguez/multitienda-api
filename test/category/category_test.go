package test

import (
	"log"
	"testing"

	"github.com/devrodriguez/multitienda-api/internal/category"
)

type categoryTestAdapter struct{}

func NewCategoryTestAdapter() category.PortOut {
	return &categoryTestAdapter{}
}

func (adapter *categoryTestAdapter) GetAllDB() ([]*category.Category, error) {
	return []*category.Category{
		{
			Name:        "Category Test",
			Description: "Esta es una categoria de prueba",
		},
	}, nil
}

func TestGetAllCategory(t *testing.T) {
	outAdapter := NewCategoryTestAdapter()
	inAdapter := category.NewAdapterIn(outAdapter)
	categories, err := inAdapter.GetAll()
	if err != nil {
		t.Error(err)
	}

	log.Printf("Categories: %d", len(categories))
}
