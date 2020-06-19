package customer

type CustomerPortOut interface {
	GetAll() ([]*Customer, error)
	Create(Customer) error
}
