package service

import (
	"product/internal/messages"
	"product/internal/repository"
)

type CategoryService interface {
	GetCategoryByID(id int64) (messages.CategoryResponse, error)
	CreateCategory(request messages.CategoryCreateRequest) error
	UpdateCategory(request messages.CategoryUpdateRequest) error
	DeleteCategory(id int64) error
}

type categoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (c categoryService) GetCategoryByID(id int64) (messages.CategoryResponse, error) {
	category, err := c.categoryRepo.GetByID(id)
	if err != nil {
		return messages.CategoryResponse{}, err
	}
	return messages.CategoryResponse{
		ID:          category.CategoryID,
		Name:        category.Name,
		Description: category.Description,
		ParentID:    category.ParentID,
	}, nil
}

func (c categoryService) GetCategoriesByParentID(parentID int64) (messages.CategoryListResponse, error) {
	categories, err := c.categoryRepo.GetChildrenByParentID(parentID)
	if err != nil {
		return messages.CategoryListResponse{}, err
	}
	var categoriesResponses []messages.CategoryResponse
	for _, category := range categories {
		categoriesResponses = append(categoriesResponses, messages.CategoryResponse{
			ID:          category.CategoryID,
			Name:        category.Name,
			Description: category.Description,
			ParentID:    category.ParentID,
		})
	}
	return messages.CategoryListResponse{
		Categories: categoriesResponses,
	}, nil
}

func (c categoryService) CreateCategory(request messages.CategoryCreateRequest) error {
	product := repository.Category{
		Name:        request.Name,
		Description: request.Description,
		ParentID:    request.ParentID,
	}

	err := c.categoryRepo.Create(&product)
	if err != nil {
		return err
	}
	return nil
}

func (c categoryService) UpdateCategory(request messages.CategoryUpdateRequest) error {
	category := repository.Category{
		CategoryID:  request.ID,
		Name:        request.Name,
		Description: request.Description,
		ParentID:    request.ParentID,
	}
	err := c.categoryRepo.Update(&category)
	if err != nil {
		return err
	}
	return nil
}

func (c categoryService) DeleteCategory(id int64) error {
	err := c.categoryRepo.DeleteByID(id)
	if err != nil {
		return err
	}
	return nil
}
