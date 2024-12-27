package repository

import "gorm.io/gorm"

type Product struct {
	// Ера. Добавь сюда поля структуры Product
}

func (Product) TableName() string {
	return "products"
}

type ProductRepository interface {
	//	Ера. Добавь сюда методы для работы с продуктами
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

//	Ера. Добавь сюда методы для работы с продуктами
