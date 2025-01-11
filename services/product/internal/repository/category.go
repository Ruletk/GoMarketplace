package repository

import (
	"github.com/Ruletk/GoMarketplace/pkg/logging"
	"gorm.io/gorm"
)

type Category struct {
	CategoryID  int64  `json:"category_id" gorm:"primaryKey" gorm:"column:category_id"`
	Name        string `json:"name" gorm:"unique" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	ParentID    int64  `json:"parent_id" gorm:"column:parent_id"`
}

func (Category) TableName() string {
	return "categories"
}

type CategoryRepository interface {
	GetByID(id int64) (*Category, error)
	GetChildrenByParentID(parentID int64) ([]*Category, error)

	Create(category *Category) error

	Update(category *Category) error

	DeleteByID(id int64) error
	DeleteByName(name string) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	logging.Logger.Info("Creating new category repository")
	return &categoryRepository{
		db: db,
	}
}

func (c categoryRepository) GetByID(id int64) (*Category, error) {
	logging.Logger.Debug("Getting category by id: ", id)
	var category Category
	err := c.db.Where("category_id = ?", id).First(&category).Error
	if err != nil {
		logging.Logger.WithError(err).Error("Error getting category by id: ", id)
		return nil, err
	}
	logging.Logger.Info("Category found: ", category)
	return &category, nil
}

func (c categoryRepository) GetChildrenByParentID(parentID int64) ([]*Category, error) {
	logging.Logger.Debug("Getting children by parent id: ", parentID)
	var categories []*Category
	err := c.db.Where("parent_id = ?", parentID).Find(&categories).Error
	if err != nil {
		logging.Logger.WithError(err).Error("Error getting children by parent id: ", parentID)
		return nil, err
	}
	logging.Logger.Info("Categories found: ", categories)
	return categories, nil
}

func (c categoryRepository) Create(category *Category) error {
	logging.Logger.Info("Creating category: ", category)
	err := c.db.Create(category).Error
	if err != nil {
		logging.Logger.WithError(err).Error("Error creating category: ", category)
		return err
	}
	logging.Logger.Info("Category created: ", category)
	return nil
}

func (c categoryRepository) Update(category *Category) error {
	logging.Logger.Debug("Updating category: ", category)
	err := c.db.Save(category).Error
	if err != nil {
		logging.Logger.WithError(err).Error("Error updating category: ", category)
		return err
	}
	logging.Logger.Info("Category updated: ", category)
	return nil
}

func (c categoryRepository) DeleteByID(id int64) error {
	logging.Logger.Debug("Deleting category by id: ", id)
	err := c.db.Where("category_id = ?", id).Delete(&Category{}).Error
	if err != nil {
		logging.Logger.WithError(err).Error("Error deleting category by id: ", id)
		return err
	}
	logging.Logger.Info("Category deleted by id: ", id)
	return nil
}

func (c categoryRepository) DeleteByName(name string) error {
	logging.Logger.Debug("Deleting category by name: ", name)
	err := c.db.Where("name = ?", name).Delete(&Category{}).Error
	if err != nil {
		logging.Logger.WithError(err).Error("Error deleting category by name: ", name)
		return err
	}
	logging.Logger.Info("Category deleted by name: ", name)
	return nil
}
