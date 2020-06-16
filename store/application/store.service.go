package store

import (
	domain "github.com/devrodriguez/multitienda-api/store/domain"
	infra "github.com/devrodriguez/multitienda-api/store/infrastructure"
)

type storeService struct {
	storeRepository infra.RepositoryContract
}

func NewStoreService(sr infra.RepositoryContract) domain.ServiceContract {
	return &storeService{
		storeRepository: sr,
	}
}

func (ss *storeService) GetAllStores() ([]*domain.Store, error) {
	return ss.storeRepository.GetAllStores()
}
