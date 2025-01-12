package service

import (
	"github.com/Ruletk/GoMarketplace/pkg/logging"
	"product/internal/messages"
	"product/internal/repository"
)

type ProductService interface {
	CreateProduct(request messages.ProductCreateRequest) error

	UpdateProduct(request messages.ProductUpdateRequest) error

	DeleteProduct(id int64) error
	DeleteAllByCompanyID(companyID int64) error

	GetProductByID(id int64) (messages.ProductResponse, error)
	GetProductsByFilter(filter messages.ProductFilter) (messages.ProductListResponse, error)
}

type productService struct {
	productRepo      repository.ProductRepository
	inventoryService InventoryService
}

func NewProductService(productRepo repository.ProductRepository, inventoryService InventoryService) ProductService {
	logging.Logger.Info("Creating new product service")
	return &productService{
		productRepo:      productRepo,
		inventoryService: inventoryService,
	}
}

func (p productService) CreateProduct(request messages.ProductCreateRequest) error {
	logging.Logger.Debug("Creating product: ", request)
	inventory, err := p.inventoryService.CreateInventory(request.Quantity)

	product := repository.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		CategoryID:  request.CategoryID,
		CompanyID:   request.CompanyID,
		InventoryID: inventory.InventoryID,
		DiscountID:  request.DiscountID,
	}
	err = p.productRepo.Create(&product)
	if err != nil {
		logging.Logger.WithError(err).Error("Error creating product: ", request)
		return err
	}
	logging.Logger.Info("Product created: ", product)
	return nil
}

func (p productService) UpdateProduct(request messages.ProductUpdateRequest) error {
	logging.Logger.Debug("Updating product: ", request)
	product := repository.Product{
		ProductID:   request.ID,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		CategoryID:  request.CategoryID,
		DiscountID:  request.DiscountID,
	}
	err := p.productRepo.Update(&product)
	if err != nil {
		logging.Logger.WithError(err).Error("Error updating product: ", request)
		return err
	}
	logging.Logger.Info("Product updated: ", product)
	return nil
}

func (p productService) DeleteProduct(id int64) error {
	logging.Logger.Debug("Deleting product by id: ", id)
	err := p.productRepo.DeleteByID(id)
	if err != nil {
		logging.Logger.WithError(err).Error("Error deleting product by id: ", id)
		return err
	}
	logging.Logger.Info("Product deleted by id: ", id)
	return nil
}

func (p productService) DeleteAllByCompanyID(companyID int64) error {
	logging.Logger.Debug("Deleting products by company id: ", companyID)
	err := p.productRepo.DeleteAllByCompanyID(companyID)
	if err != nil {
		logging.Logger.WithError(err).Error("Error deleting products by company id: ", companyID)
		return err
	}
	logging.Logger.Info("Products deleted by company id: ", companyID)
	return nil
}

func (p productService) GetProductByID(id int64) (messages.ProductResponse, error) {
	logging.Logger.Debug("Getting product by id: ", id)
	product, err := p.productRepo.GetByID(id)
	if err != nil {
		logging.Logger.WithError(err).Error("Error getting product by id: ", id)
		return messages.ProductResponse{}, err
	}
	logging.Logger.Info("Product found: ", product)
	return messages.ProductResponse{
		ID:          product.ProductID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CategoryID:  product.CategoryID,
		CompanyID:   product.CompanyID,
		InventoryID: product.InventoryID,
		DiscountID:  product.DiscountID,
	}, nil
}

func (p productService) GetProductsByFilter(filter messages.ProductFilter) (messages.ProductListResponse, error) {
	logging.Logger.Debug("Getting products by filter: ", filter)
	validateFilter(&filter)

	products, totalCount, err := p.productRepo.GetByFilter(&filter)
	if err != nil {
		logging.Logger.WithError(err).Error("Error getting products by filter: ", filter)
		return messages.ProductListResponse{}, err
	}
	logging.Logger.Info("Products found: ", products)
	productResponses := productsResponseFromModels(products)

	return messages.ProductListResponse{
		Products:   productResponses,
		TotalCount: totalCount,
	}, nil
}

func productResponseFromModel(p *repository.Product) messages.ProductResponse {
	logging.Logger.Debug("Creating product response from model: ", p)
	return messages.ProductResponse{
		ID:          p.ProductID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		CategoryID:  p.CategoryID,
		CompanyID:   p.CompanyID,
		InventoryID: p.InventoryID,
		DiscountID:  p.DiscountID,
	}
}

func productsResponseFromModels(products []*repository.Product) []messages.ProductResponse {
	logging.Logger.Debug("Creating products response from models: ", products)
	productResponses := make([]messages.ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = productResponseFromModel(product)
	}
	return productResponses
}

func validateFilter(filter *messages.ProductFilter) {
	logging.Logger.Debug("Validating filter: ", filter)
	if filter.PageSize <= 0 || filter.PageSize > 100 {
		filter.PageSize = 10
	}
	if filter.PageNumber < 0 {
		filter.PageNumber = 0
	}
}
