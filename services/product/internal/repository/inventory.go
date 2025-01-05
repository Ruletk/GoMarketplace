package repository

import "gorm.io/gorm"

type Inventory struct {
	InventoryID int64 `json:"inventory_id" gorm:"primaryKey" gorm:"column:inventory_id"`
	Quantity    int64 `json:"quantity" gorm:"column:quantity"`
}

func (Inventory) TableName() string {
	return "inventories"
}

type InventoryRepository interface {
	GetByID(id int64) (*Inventory, error)

	Create(inventory *Inventory) error

	Update(inventory *Inventory) error

	DeleteByID(id int64) error
}

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepository{
		db: db,
	}
}

func (i inventoryRepository) GetByID(id int64) (*Inventory, error) {
	var inventory Inventory
	err := i.db.Where("inventory_id = ?", id).First(&inventory).Error
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

func (i inventoryRepository) Create(inventory *Inventory) error {
	return i.db.Create(inventory).Error
}

func (i inventoryRepository) Update(inventory *Inventory) error {
	return i.db.Save(inventory).Error
}

func (i inventoryRepository) DeleteByID(id int64) error {
	return i.db.Where("inventory_id = ?", id).Delete(&Inventory{}).Error
}
