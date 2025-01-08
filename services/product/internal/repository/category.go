package repository

import "gorm.io/gorm"

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
	return &categoryRepository{
		db: db,
	}
}

func (c categoryRepository) GetByID(id int64) (*Category, error) {
	var category Category
	err := c.db.Where("category_id = ?", id).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (c categoryRepository) GetChildrenByParentID(parentID int64) ([]*Category, error) {
	var categories []*Category
	err := c.db.Where("parent_id = ?", parentID).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (c categoryRepository) Create(category *Category) error {
	return c.db.Create(category).Error

}

func (c categoryRepository) Update(category *Category) error {
	return c.db.Save(category).Error

}

func (c categoryRepository) DeleteByID(id int64) error {
	return c.db.Where("category_id = ?", id).Delete(&Category{}).Error
}

func (c categoryRepository) DeleteByName(name string) error {
	return c.db.Where("name = ?", name).Delete(&Category{}).Error
}
