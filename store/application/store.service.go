package store

import (
	domain "github.com/devrodriguez/multitienda-api/store/domain"
	infra "github.com/devrodriguez/multitienda-api/store/infrastructure"
)

type storeService struct {
	storeRepository infra.RepositoryContract
}

// NewStoreService return a new store service interface
func NewStoreService(sr infra.RepositoryContract) domain.ServiceContract {
	return &storeService{
		storeRepository: sr,
	}
}

// GetAllStores implements the bussines logic
func (ss *storeService) GetAllStores() ([]*domain.Store, error) {
	var spStore []*domain.Store
	stores, err := ss.storeRepository.GetAllStores()
	if err != nil {
		return []*domain.Store{}, err
	}

	// Logic implementation example
	for _, store := range stores {
		if store.Name == "El Negocio" {
			store.Name += "@"
			spStore = append(spStore, store)
			break
		}
	}
	return spStore, nil
}
