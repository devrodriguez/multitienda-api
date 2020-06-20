package customer

import (
	"log"
	"testing"

	"github.com/devrodriguez/multitienda-api/internal/customer"
)

type testCustomerAdapter struct{}

func NewTestAdapter() customer.PortOut {
	return &testCustomerAdapter{}
}

func (tc *testCustomerAdapter) GetAllDB() ([]*customer.Customer, error) {
	return []*customer.Customer{
		{
			Name: "John Rodriguez",
		},
	}, nil
}

func (tc *testCustomerAdapter) CreateDB(customer customer.Customer) error {
	log.Println("Customer created")
	return nil
}
func TestCustomer(t *testing.T) {
	outAdapter := NewTestAdapter()
	inAdapter := customer.NewAdapterIn(outAdapter)
	stores, err := inAdapter.GetAll()
	if err != nil {
		t.Error(err)
	}

	log.Printf("Customers: %d", len(stores))
}
