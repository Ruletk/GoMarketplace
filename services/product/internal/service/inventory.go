package service

import (
	"product/internal/messages"
	"product/internal/repository"
)

type InventoryService interface {
	CreateInventory(request messages.InventoryCreateRequest) (int, error)

	UpdateInventory(request messages.InventoryUpdateRequest) error

	DeleteInventory(request messages.InventoryDeleteRequest) error
}

type inventoryService struct {
	inventoryRepo repository.InventoryRepository
}

func NewInventoryService(inventoryRepo repository.InventoryRepository) InventoryService {
	return &inventoryService{
		inventoryRepo: inventoryRepo,
	}
}

func (i inventoryService) CreateInventory(request messages.InventoryCreateRequest) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (i inventoryService) UpdateInventory(request messages.InventoryUpdateRequest) error {
	//TODO implement me
	panic("implement me")
}

func (i inventoryService) DeleteInventory(request messages.InventoryDeleteRequest) error {
	//TODO implement me
	panic("implement me")
}
