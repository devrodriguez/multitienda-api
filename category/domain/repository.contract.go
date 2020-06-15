package category

type RepositoryContract interface {
	GetAll() ([]*Category, error)
}
