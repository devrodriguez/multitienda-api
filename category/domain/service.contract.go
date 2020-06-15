package category

type ServiceContract interface {
	GetAll() ([]*Category, error)
}
