package service

import (
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
	return &inventoryService{
		inventoryRepo: inventoryRepo,
	}
}

func (i inventoryService) GetInventoryByID(id int64) (messages.InventoryResponse, error) {
	inventory, err := i.inventoryRepo.GetByID(id)
	if err != nil {
		return messages.InventoryResponse{}, err
	}
	return messages.InventoryResponse{
		ID:       inventory.InventoryID,
		Quantity: inventory.Quantity,
	}, nil
}

func (i inventoryService) CreateInventory(quantity int64) (*repository.Inventory, error) {
	inv := repository.Inventory{
		Quantity: quantity,
	}
	inventory, err := i.inventoryRepo.Create(&inv)
	if err != nil {
		// TODO: log error
		return nil, err
	}
	return inventory, nil
}

func (i inventoryService) UpdateInventory(inventory *repository.Inventory) error {
	return i.inventoryRepo.Update(inventory)
}

func (i inventoryService) DeleteInventory(inventory *repository.Inventory) error {
	return i.inventoryRepo.DeleteByID(inventory.InventoryID)
}
