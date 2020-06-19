package store

type ServiceContract interface {
	GetAllStores() ([]*Store, error)
}
