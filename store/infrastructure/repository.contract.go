package store

import domain "github.com/devrodriguez/multitienda-api/store/domain"

type RepositoryContract interface {
	GetAllStores() ([]*domain.Store, error)
}
