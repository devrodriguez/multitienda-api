package category

type PortOut interface {
	GetAllDB() ([]*Category, error)
}
