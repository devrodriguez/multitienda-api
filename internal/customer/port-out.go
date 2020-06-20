package customer

type PortOut interface {
	GetAllDB() ([]*Customer, error)
	CreateDB(Customer) error
}
