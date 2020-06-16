package category

import (
	domain "github.com/devrodriguez/multitienda-api/category/domain"
	infra "github.com/devrodriguez/multitienda-api/category/infrastructure"
)

type categoryService struct {
	categoryRepository infra.RepositoryContract
}

func NewCategoryService(repositoryContract infra.RepositoryContract) domain.ServiceContract {
	return &categoryService{
		repositoryContract,
	}
}

func (cs *categoryService) GetAll() ([]*domain.Category, error) {
	return cs.categoryRepository.GetAll()
}
