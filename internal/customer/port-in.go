package customer

type PortIn interface {
	GetAll() ([]*Customer, error)
	Create(Customer) error
}
