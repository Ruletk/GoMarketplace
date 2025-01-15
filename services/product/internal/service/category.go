package service

import (
	"github.com/Ruletk/GoMarketplace/pkg/logging"
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
	logging.Logger.Info("Creating new category service")
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

func (c categoryService) GetCategoryByID(id int64) (messages.CategoryResponse, error) {
	logging.Logger.Debug("Getting category by id: ", id)

	category, err := c.categoryRepo.GetByID(id)
	if err != nil {
		logging.Logger.WithError(err).Error("Error getting category by id: ", id)
		return messages.CategoryResponse{}, err
	}
	logging.Logger.Info("Category found: ", category)
	return messages.CategoryResponse{
		ID:          category.CategoryID,
		Name:        category.Name,
		Description: category.Description,
		ParentID:    category.ParentID,
	}, nil
}

func (c categoryService) GetCategoriesByParentID(parentID int64) (messages.CategoryListResponse, error) {
	logging.Logger.Debug("Getting children by parent id: ", parentID)
	categories, err := c.categoryRepo.GetChildrenByParentID(parentID)
	if err != nil {
		logging.Logger.WithError(err).Error("Error getting children by parent id: ", parentID)
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
	logging.Logger.Info("Categories found: ", categoriesResponses)
	return messages.CategoryListResponse{
		Categories: categoriesResponses,
	}, nil
}

func (c categoryService) CreateCategory(request messages.CategoryCreateRequest) error {
	logging.Logger.Debug("Creating category with name: ", request.Name)
	product := repository.Category{
		Name:        request.Name,
		Description: request.Description,
		ParentID:    request.ParentID,
	}

	err := c.categoryRepo.Create(&product)
	if err != nil {
		logging.Logger.WithError(err).Error("Error creating category with name: ", request.Name)
		return err
	}
	logging.Logger.Info("Category created with name: ", request.Name)
	return nil
}

func (c categoryService) UpdateCategory(request messages.CategoryUpdateRequest) error {
	logging.Logger.Debug("Updating category with id: ", request.ID)
	category := repository.Category{
		CategoryID:  request.ID,
		Name:        request.Name,
		Description: request.Description,
		ParentID:    request.ParentID,
	}
	err := c.categoryRepo.Update(&category)
	if err != nil {
		logging.Logger.WithError(err).Error("Error updating category with id: ", request.ID)
		return err
	}
	logging.Logger.Info("Category updated with id: ", request.ID)
	return nil
}

func (c categoryService) DeleteCategory(id int64) error {
	logging.Logger.Debug("Deleting category by id: ", id)
	err := c.categoryRepo.DeleteByID(id)
	if err != nil {
		logging.Logger.WithError(err).Error("Error deleting category by id: ", id)
		return err
	}
	logging.Logger.Info("Category deleted by id: ", id)
	return nil
}
