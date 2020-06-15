package category

import category "github.com/devrodriguez/multitienda-api/category/domain"

type categoryService struct {
	categoryRepository category.RepositoryContract
}

func NewCategoryService(repositoryContract category.RepositoryContract) category.ServiceContract {
	return &categoryService{
		repositoryContract,
	}
}

func (rs *categoryService) GetAll() ([]*category.Category, error) {
	return rs.categoryRepository.GetAll()
}
