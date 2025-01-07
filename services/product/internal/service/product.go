package service

import (
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
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (p productService) CreateProduct(request messages.ProductCreateRequest) error {
	product := repository.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		CategoryID:  request.CategoryID,
		CompanyID:   request.CompanyID,
		InventoryID: request.InventoryID,
		DiscountID:  request.DiscountID,
	}
	//TODO: log product creation
	err := p.productRepo.Create(&product)
	if err != nil {
		//TODO: log error
		return err
	}
	//TODO: log product creation success
	return nil
}

func (p productService) UpdateProduct(request messages.ProductUpdateRequest) error {
	product := repository.Product{
		ProductID:   request.ID,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		CategoryID:  request.CategoryID,
		DiscountID:  request.DiscountID,
	}
	//TODO: log product update
	err := p.productRepo.Update(&product)
	if err != nil {
		//TODO: log error
		return err
	}
	//TODO: log product update success
	return nil
}

func (p productService) DeleteProduct(id int64) error {
	//TODO: log product deletion
	err := p.productRepo.DeleteByID(id)
	if err != nil {
		//TODO: log error
		return err
	}
	//TODO: log product deletion success
	return nil
}

func (p productService) DeleteAllByCompanyID(companyID int64) error {
	//TODO: log product deletion
	err := p.productRepo.DeleteAllByCompanyID(companyID)
	if err != nil {
		//TODO: log error
		return err
	}
	//TODO: log product deletion success
	return nil
}

func (p productService) GetProductByID(id int64) (messages.ProductResponse, error) {
	//TODO: log product retrieval
	product, err := p.productRepo.GetByID(id)
	if err != nil {
		//TODO: log error
		return messages.ProductResponse{}, err
	}
	//TODO: log product retrieval success
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
	//TODO: log product retrieval
	validateFilter(&filter)

	products, err := p.productRepo.GetByFilter(&filter)
	if err != nil {
		//TODO: log error
		return messages.ProductListResponse{}, err
	}
	//TODO: log product retrieval success
	productResponses := productsResponseFromModels(products)

	return messages.ProductListResponse{
		Products: productResponses,
	}, nil
}

func productResponseFromModel(p *repository.Product) messages.ProductResponse {
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
	productResponses := make([]messages.ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = productResponseFromModel(product)
	}
	return productResponses
}

func validateFilter(filter *messages.ProductFilter) {
	if filter.PageSize <= 0 || filter.PageSize > 100 {
		filter.PageSize = 10
	}
	if filter.PageNumber < 0 {
		filter.PageNumber = 0
	}
}
