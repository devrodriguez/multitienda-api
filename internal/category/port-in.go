package category

type PortIn interface {
	GetAll() ([]*Category, error)
}
