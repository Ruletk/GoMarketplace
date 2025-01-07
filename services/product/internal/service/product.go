package service

import (
	"product/internal/messages"
	"product/internal/repository"
)

type ProductService interface {
	CreateProduct(request messages.ProductCreateRequest) (int, error)
	BatchCreateProduct(request messages.ProductBatchCreateRequest) ([]int, error)

	UpdateProduct(request messages.ProductUpdateRequest) error
	BatchUpdateProduct(request messages.ProductBatchUpdateRequest) error

	DeleteProduct(request messages.ProductDeleteRequest) error
	BatchDeleteProduct(request messages.ProductBatchDeleteRequest) error
	DeleteAllByCompanyID(companyID int) error

	GetProductByID(id int) (messages.ProductResponse, error)
	GetProductByFilter(filter messages.ProductFilter) ([]messages.ProductResponse, error)
	GetProductsByCategory(categoryID int) ([]messages.ProductResponse, error)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (p productService) CreateProduct(request messages.ProductCreateRequest) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (p productService) BatchCreateProduct(request messages.ProductBatchCreateRequest) ([]int, error) {
	//TODO implement me
	panic("implement me")
}

func (p productService) UpdateProduct(request messages.ProductUpdateRequest) error {
	//TODO implement me
	panic("implement me")
}

func (p productService) BatchUpdateProduct(request messages.ProductBatchUpdateRequest) error {
	//TODO implement me
	panic("implement me")
}

func (p productService) DeleteProduct(request messages.ProductDeleteRequest) error {
	//TODO implement me
	panic("implement me")
}

func (p productService) BatchDeleteProduct(request messages.ProductBatchDeleteRequest) error {
	//TODO implement me
	panic("implement me")
}

func (p productService) DeleteAllByCompanyID(companyID int) error {
	//TODO implement me
	panic("implement me")
}

func (p productService) GetProductByID(id int) (messages.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p productService) GetProductByFilter(filter messages.ProductFilter) ([]messages.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p productService) GetProductsByCategory(categoryID int) ([]messages.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}
