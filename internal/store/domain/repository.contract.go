package store

type RepositoryContract interface {
	GetAllStores() ([]*Store, error)
}
