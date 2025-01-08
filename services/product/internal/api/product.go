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

func NewProductAPI(
	inventoryService service.InventoryService,
	productService service.ProductService,
	categoryService service.CategoryService,
) *ProductAPI {
	return &ProductAPI{
		inventoryService: inventoryService,
		productService:   productService,
		categoryService:  categoryService,
	}
}

func (api *ProductAPI) RegisterPublicOnlyRoutes(router *gin.RouterGroup) {}

func (api *ProductAPI) RegisterPublicRoutes(router *gin.RouterGroup) {
	//TODO: Add logging
	router.GET("/products", api.GetProducts)
	router.GET("/:id", api.GetProductByID)
	router.GET("/category/:id", api.GetCategoryByID)
	router.GET("/category/:id/products", api.GetProductsByCategory)
}

func (api *ProductAPI) RegisterPrivateRoutes(router *gin.RouterGroup) {
	//TODO: Add logging
	router.POST("/products", api.CreateProduct)
	router.PUT("/products", api.UpdateProduct)
	router.DELETE("/:id", api.DeleteProduct)

	router.POST("/categories", api.CreateCategory)
	router.PUT("/categories", api.UpdateCategory)
	router.DELETE("/category/:id", api.DeleteCategory)
}

// Handlers

func (api *ProductAPI) GetProductByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if id <= 0 || err != nil {
		//TODO: Add logging
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid product ID",
		})
		return
	}
	//TODO: Add logging
	response, err := api.productService.GetProductByID(id)
	if err != nil {
		//TODO: log error
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to fetch product",
		})
		return
	}
	//TODO: log success
	c.JSON(http.StatusOK, response)
}

func (api *ProductAPI) GetProducts(c *gin.Context) {
	var queryParams ProductQueryParams
	queryParams.MaxPrice = math.MaxFloat64
	//TODO: Add logging

	if err := c.ShouldBindQuery(&queryParams); err != nil {
		//TODO: log error
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid query parameters",
		})
		return
	}

	filter := NewFilterFromQueryParams(queryParams)
	//TODO: log success, debug log params

	products, err := api.productService.GetProductsByFilter(filter)
	if err != nil {
		//TODO: log error
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to fetch products",
		})
		return
	}
	//TODO: log success

	c.JSON(http.StatusOK, products)
}

func (api *ProductAPI) CreateProduct(c *gin.Context) {
	var req messages.ProductCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		//TODO: log error
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}
	//TODO: log creation request

	if err := api.productService.CreateProduct(req); err != nil {
		//TODO: log error
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to create product",
		})
		return
	}
	//TODO: log success

	c.JSON(http.StatusCreated, messages.ApiResponse{
		Code:    http.StatusCreated,
		Type:    "success",
		Message: "Product created successfully",
	})
}

func (api *ProductAPI) UpdateProduct(c *gin.Context) {
	var req messages.ProductUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		//TODO: log error
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}
	//TODO: log update request

	if err := api.productService.UpdateProduct(req); err != nil {
		//TODO: log error
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to update product",
		})
		return
	}

	//TODO: log success

	c.JSON(http.StatusOK, messages.ApiResponse{
		Code:    http.StatusOK,
		Type:    "success",
		Message: "Product updated successfully",
	})
}

func (api *ProductAPI) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if id <= 0 || err != nil {
		//TODO: log error
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid product ID",
		})
		return
	}
	//TODO: log delete request
	if err := api.productService.DeleteProduct(int64(id)); err != nil {
		//TODO: log error
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to delete product",
		})
		return
	}

	//TODO: log success

	c.JSON(http.StatusOK, messages.ApiResponse{
		Code:    http.StatusOK,
		Type:    "success",
		Message: "Product deleted successfully",
	})
}

func (api *ProductAPI) GetCategoryByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if id <= 0 || err != nil {
		//TODO: Add logging
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid category ID",
		})
		return
	}

	response, err := api.categoryService.GetCategoryByID(id)
	//TODO: log request
	if err != nil {
		//TODO: log error
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to fetch category",
		})
		return
	}

	//TODO: log success
	c.JSON(http.StatusOK, response)
}

func (api *ProductAPI) GetProductsByCategory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if id <= 0 || err != nil {
		//TODO: Add logging
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid category ID",
		})
		return
	}
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
		return
	}

	//TODO: log success

	c.JSON(http.StatusOK, products)
}

func (api *ProductAPI) CreateCategory(c *gin.Context) {
	var req messages.CategoryCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		//TODO: log error
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}

	err := api.categoryService.CreateCategory(req)
	if err != nil {
		//TODO: log error
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to create category",
		})
		return
	}

	c.JSON(http.StatusCreated, messages.ApiResponse{
		Code:    http.StatusCreated,
		Type:    "success",
		Message: "Category created successfully",
	})
}

func (api *ProductAPI) UpdateCategory(c *gin.Context) {
	var req messages.CategoryUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		//TODO: log error
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}

	err := api.categoryService.UpdateCategory(req)
	if err != nil {
		//TODO: log error
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to update category",
		})
		return
	}

	//TODO: log success
	c.JSON(http.StatusOK, messages.ApiResponse{
		Code:    http.StatusOK,
		Type:    "success",
		Message: "Category updated successfully",
	})
}

func (api *ProductAPI) DeleteCategory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if id <= 0 || err != nil {
		//TODO: log error
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid category ID",
		})
		return
	}

	err = api.categoryService.DeleteCategory(id)
	if err != nil {
		//TODO: log error
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to delete category",
		})
		return
	}

	//TODO: log success
	c.JSON(http.StatusOK, messages.ApiResponse{
		Code:    http.StatusOK,
		Type:    "success",
		Message: "Category deleted successfully",
	})
}
