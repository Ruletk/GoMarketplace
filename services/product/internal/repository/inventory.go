package repository

import (
	"github.com/Ruletk/GoMarketplace/pkg/logging"
	"gorm.io/gorm"
)

type Inventory struct {
	InventoryID int64 `json:"inventory_id" gorm:"primaryKey" gorm:"column:inventory_id"`
	Quantity    int64 `json:"quantity" gorm:"column:quantity"`
}

func (Inventory) TableName() string {
	return "inventories"
}

type InventoryRepository interface {
	GetByID(id int64) (*Inventory, error)

	Create(inventory *Inventory) (*Inventory, error)

	Update(inventory *Inventory) error

	DeleteByID(id int64) error
}

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	logging.Logger.Info("Creating new inventory repository")
	return &inventoryRepository{
		db: db,
	}
}

func (i inventoryRepository) GetByID(id int64) (*Inventory, error) {
	logging.Logger.Debug("Getting inventory by id: ", id)
	var inventory Inventory
	if err := i.db.Where("inventory_id = ?", id).First(&inventory).Error; err != nil {
		logging.Logger.WithError(err).Error("Error getting inventory by id: ", id)
		return nil, err
	}
	logging.Logger.Info("Inventory found: ", inventory)
	return &inventory, nil
}

func (i inventoryRepository) Create(inventory *Inventory) (*Inventory, error) {
	logging.Logger.Debug("Creating inventory: ", inventory)
	if err := i.db.Create(inventory).Error; err != nil {
		logging.Logger.WithError(err).Error("Error creating inventory: ", inventory)
		return nil, err
	}
	logging.Logger.Info("Inventory created: ", inventory)
	return inventory, nil
}

func (i inventoryRepository) Update(inventory *Inventory) error {
	logging.Logger.Debug("Updating inventory: ", inventory)
	if err := i.db.Save(inventory).Error; err != nil {
		logging.Logger.WithError(err).Error("Error updating inventory: ", inventory)
		return err
	}
	logging.Logger.Info("Inventory updated: ", inventory)
	return nil
}

func (i inventoryRepository) DeleteByID(id int64) error {
	logging.Logger.Debug("Deleting inventory by id: ", id)

	if err := i.db.Where("inventory_id = ?", id).Delete(&Inventory{}).Error; err != nil {
		logging.Logger.WithError(err).Error("Error deleting inventory by id: ", id)
		return err
	}
	logging.Logger.Info("Inventory deleted by id: ", id)
	return nil
}
