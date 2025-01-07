package service

import (
	"product/internal/messages"
	"product/internal/repository"
)

type CategoryService interface {
	CreateCategory(request messages.CategoryCreateRequest) (int, error)
	BatchCreateCategory(request messages.CategoryBatchCreateRequest) ([]int, error)

	UpdateCategory(request messages.CategoryUpdateRequest) error
	BatchUpdateCategory(request messages.CategoryBatchUpdateRequest) error

	DeleteCategory(request messages.CategoryDeleteRequest) error
	BatchDeleteCategory(request messages.CategoryBatchDeleteRequest) error
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (c categoryService) CreateCategory(request messages.CategoryCreateRequest) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (c categoryService) BatchCreateCategory(request messages.CategoryBatchCreateRequest) ([]int, error) {
	//TODO implement me
	panic("implement me")
}

func (c categoryService) UpdateCategory(request messages.CategoryUpdateRequest) error {
	//TODO implement me
	panic("implement me")
}

func (c categoryService) BatchUpdateCategory(request messages.CategoryBatchUpdateRequest) error {
	//TODO implement me
	panic("implement me")
}

func (c categoryService) DeleteCategory(request messages.CategoryDeleteRequest) error {
	//TODO implement me
	panic("implement me")
}

func (c categoryService) BatchDeleteCategory(request messages.CategoryBatchDeleteRequest) error {
	//TODO implement me
	panic("implement me")
}
