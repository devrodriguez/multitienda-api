package store

import (
	"testing"

	sa "github.com/devrodriguez/multitienda-api/internal/store/application"
	sd "github.com/devrodriguez/multitienda-api/internal/store/domain"
	si "github.com/devrodriguez/multitienda-api/internal/store/infrastructure"
	"github.com/stretchr/testify/assert"
)

type testRepository struct{}

func NewTestRepository() si.RepositoryContract {
	return &testRepository{}
}

func (tr *testRepository) GetAllStores() ([]*sd.Store, error) {
	return []*sd.Store{
		{
			Name: "El Negocio",
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

	assert.Equal(t, s[0].Name, "El Negocio@", "Comparison error")
}

func TestStoresCount(t *testing.T) {
	sr := NewTestRepository()
	ss := sa.NewStoreService(sr)
	s, err := ss.GetAllStores()
	if err != nil {
		t.Error(err)
	}

	assert.Greater(t, len(s), 0, "Returned items less or equal 0")
}
