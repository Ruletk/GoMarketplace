package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"product/internal/messages"
	"product/internal/service"
	"strconv"
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
	router.GET("/product/:id", api.GetProductByID)
	router.GET("/category/:id", api.GetCategoryByID)
	router.GET("/category/:id/product", api.GetProductsByCategory)
}

func (api *ProductAPI) RegisterPublicRoutes(router *gin.RouterGroup) {
	router.GET("/products", api.GetProducts)
}

func (api *ProductAPI) RegisterPrivateRoutes(router *gin.RouterGroup) {
	router.POST("/products", api.CreateProduct)
	router.PUT("/products", api.UpdateProduct)
	router.DELETE("/product/:id", api.DeleteProduct)

	router.POST("/categories", api.CreateCategory)
	router.PUT("/categories", api.UpdateCategory)
	router.DELETE("/category/:id", api.DeleteCategory)
}

// Handlers

func (api *ProductAPI) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	// Add logic to fetch product by ID.
	c.JSON(http.StatusOK, gin.H{"id": id, "message": "GetProductByID not implemented yet"})
}

func (api *ProductAPI) GetProducts(c *gin.Context) {
	var queryParams ProductQueryParams
	queryParams.MaxPrice = math.MaxFloat64

	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid query parameters",
		})
		return
	}

	filter := NewFilterFromQueryParams(queryParams)

	products, err := api.productService.GetProductsByFilter(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to fetch products",
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (api *ProductAPI) CreateProduct(c *gin.Context) {
	var req messages.ProductCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}

	if err := api.productService.CreateProduct(req); err != nil {
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to create product",
		})
		return
	}

	c.JSON(http.StatusOK, messages.ApiResponse{
		Code:    http.StatusOK,
		Type:    "success",
		Message: "Product created successfully",
	})
}

func (api *ProductAPI) UpdateProduct(c *gin.Context) {
	var req messages.ProductUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}

	if err := api.productService.UpdateProduct(req); err != nil {
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to update product",
		})
		return
	}

	c.JSON(http.StatusOK, messages.ApiResponse{
		Code:    http.StatusOK,
		Type:    "success",
		Message: "Product updated successfully",
	})
}

func (api *ProductAPI) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := api.productService.DeleteProduct(int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to delete product",
		})
		return
	}

	c.JSON(http.StatusOK, messages.ApiResponse{
		Code:    http.StatusOK,
		Type:    "success",
		Message: "Product deleted successfully",
	})
}

func (api *ProductAPI) GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	// Add logic to fetch category by ID.
	c.JSON(http.StatusOK, gin.H{"id": id, "message": "GetCategoryByID not implemented yet"})
}

func (api *ProductAPI) GetProductsByCategory(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	products, err := api.productService.GetProductsByFilter(messages.ProductFilter{
		CategoryIDs: []int64{id},
	})
	if err != nil {
		//TODO: log error
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to fetch products",
		})
	}

	c.JSON(http.StatusOK, products)
}

func (api *ProductAPI) CreateCategory(c *gin.Context) {
	// Add logic to create a category.
	c.JSON(http.StatusOK, gin.H{"message": "CreateCategory not implemented yet"})
}

func (api *ProductAPI) UpdateCategory(c *gin.Context) {
	// Add logic to update a category.
	c.JSON(http.StatusOK, gin.H{"message": "UpdateCategory not implemented yet"})
}

func (api *ProductAPI) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	// Add logic to delete a category by ID.
	c.JSON(http.StatusOK, gin.H{"id": id, "message": "DeleteCategory not implemented yet"})
}
