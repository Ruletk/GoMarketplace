package api

import (
	"github.com/gin-gonic/gin"
	"product/internal/service"
)

type ProductAPI struct {
	inventoryService service.InventoryService
	productService   service.ProductService
	categoryService  service.CategoryService
}

func NewProductAPI(inventoryService service.InventoryService, productService service.ProductService, categoryService service.CategoryService) *ProductAPI {
	return &ProductAPI{inventoryService: inventoryService, productService: productService, categoryService: categoryService}
}

func (api *ProductAPI) RegisterPublicOnlyRoutes(router *gin.RouterGroup) {
	// Get product here.
}

func (api *ProductAPI) RegisterPublicRoutes(router *gin.RouterGroup) {
	// Get all products here.
}

func (api *ProductAPI) RegisterPrivateRoutes(router *gin.RouterGroup) {
	// Create, update and delete product here.
}
