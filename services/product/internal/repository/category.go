package repository

import "gorm.io/gorm"

type Category struct {
	//	Ера. Сюда добавить поля для категории
}

func (Category) TableName() string {
	return "categories"
}

type CategoryRepository interface {
	//	Ера. Сюда добавить методы для работы с категориями
	//	Например, создание, получение, обновление и удаление категории
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

//	Ера. Сюда добавить методы для работы с категориями
