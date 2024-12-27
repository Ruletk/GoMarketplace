package repository

import "gorm.io/gorm"

type Inventory struct {
	//	Ера. Добавь сюда поля из картинки
}

func (Inventory) TableName() string {
	return "inventories"
}

type InventoryRepository interface {
	//	Ера. Добавь сюда методы для работы с инвентарем
}

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepository{
		db: db,
	}
}

//	Ера. Добавь сюда методы для работы с инвентарем
