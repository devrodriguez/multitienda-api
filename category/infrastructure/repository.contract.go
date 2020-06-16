package category

import domain "github.com/devrodriguez/multitienda-api/category/domain"

type RepositoryContract interface {
	GetAll() ([]*domain.Category, error)
}
