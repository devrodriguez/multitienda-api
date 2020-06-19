package customer

type CustomerPortIn interface {
	GetAll() ([]*Customer, error)
	Create(Customer) error
}
