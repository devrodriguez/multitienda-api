package store

import (
	"log"
	"testing"

	sa "github.com/devrodriguez/multitienda-api/store/application"
	sd "github.com/devrodriguez/multitienda-api/store/domain"
)

type testRepository struct{}

func NewTestRepository() RepositoryContract {
	return &testRepository{}
}

func (tr *testRepository) GetAllStores() ([]*sd.Store, error) {
	return []*sd.Store{
		{
			Name: "StoreTest",
		},
	}, nil
}

func TestGetStores(t *testing.T) {
	sr := NewTestRepository()
	ss := sa.NewStoreService(sr)
	s, err := ss.GetAllStores()
	if err != nil {
		t.Error(err)
	}

	log.Print(s)
}
