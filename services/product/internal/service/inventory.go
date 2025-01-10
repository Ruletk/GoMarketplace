package service

import (
	"github.com/Ruletk/GoMarketplace/pkg/logging"
	"product/internal/messages"
	"product/internal/repository"
)

type InventoryService interface {
	GetInventoryByID(id int64) (messages.InventoryResponse, error)

	CreateInventory(quantity int64) (*repository.Inventory, error)

	UpdateInventory(inventory *repository.Inventory) error

	DeleteInventory(inventory *repository.Inventory) error
}

type inventoryService struct {
	inventoryRepo repository.InventoryRepository
}

func NewInventoryService(inventoryRepo repository.InventoryRepository) InventoryService {
	logging.Logger.Info("Creating new inventory service")
	return &inventoryService{
		inventoryRepo: inventoryRepo,
	}
}

func (i inventoryService) GetInventoryByID(id int64) (messages.InventoryResponse, error) {
	logging.Logger.Debug("Getting inventory by id: ", id)
	inventory, err := i.inventoryRepo.GetByID(id)
	if err != nil {
		logging.Logger.WithError(err).Error("Error getting inventory by id: ", id)
		return messages.InventoryResponse{}, err
	}
	logging.Logger.Info("Inventory found: ", inventory)
	return messages.InventoryResponse{
		ID:       inventory.InventoryID,
		Quantity: inventory.Quantity,
	}, nil
}

func (i inventoryService) CreateInventory(quantity int64) (*repository.Inventory, error) {
	logging.Logger.Debug("Creating inventory with quantity: ", quantity)
	inv := repository.Inventory{
		Quantity: quantity,
	}
	inventory, err := i.inventoryRepo.Create(&inv)
	if err != nil {
		logging.Logger.WithError(err).Error("Error creating inventory with quantity: ", quantity)
		return nil, err
	}
	return inventory, nil
}

func (i inventoryService) UpdateInventory(inventory *repository.Inventory) error {
	logging.Logger.Debug("Updating inventory: ", inventory)
	if err := i.inventoryRepo.Update(inventory); err != nil {
		logging.Logger.WithError(err).Error("Error updating inventory: ", inventory)
		return err
	}
	logging.Logger.Info("Inventory updated: ", inventory)
	return nil
}

func (i inventoryService) DeleteInventory(inventory *repository.Inventory) error {
	logging.Logger.Debug("Deleting inventory: ", inventory)
	if err := i.inventoryRepo.DeleteByID(inventory.InventoryID); err != nil {
		logging.Logger.WithError(err).Error("Error deleting inventory: ", inventory)
		return err
	}
	logging.Logger.Info("Inventory deleted: ", inventory)
	return nil
}
