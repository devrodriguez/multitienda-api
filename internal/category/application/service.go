package category

import (
	domain "github.com/devrodriguez/multitienda-api/internal/category/domain"
)

type categoryService struct {
	categoryRepository domain.RepositoryContract
}

func NewCategoryService(repositoryContract domain.RepositoryContract) domain.ServiceContract {
	return &categoryService{
		repositoryContract,
	}
}

func (cs *categoryService) GetAll() ([]*domain.Category, error) {
	return cs.categoryRepository.GetAll()
}
