package api

import (
	"fmt"
	"github.com/Ruletk/GoMarketplace/pkg/logging"
	"github.com/gin-gonic/gin"
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
	logging.Logger.Info("Registering public routes")
	router.GET("/products", api.GetProducts)
	router.GET("/:id", api.GetProductByID)
	router.GET("/category/:id", api.GetCategoryByID)
	router.GET("/category/:id/products", api.GetProductsByCategory)
}

func (api *ProductAPI) RegisterPrivateRoutes(router *gin.RouterGroup) {
	logging.Logger.Info("Registering private routes")
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
		logging.Logger.WithError(err).Error("Invalid product ID")
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid product ID",
		})
		return
	}
	logging.Logger.Info("Fetching product by ID, ID: ", id)
	response, err := api.productService.GetProductByID(id)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to fetch product")
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to fetch product",
		})
		return
	}
	logging.Logger.Info("Product fetched successfully")
	c.JSON(http.StatusOK, response)
}

func (api *ProductAPI) GetProducts(c *gin.Context) {
	var queryParams ProductQueryParams
	logging.Logger.Info("Fetching products")

	if err := c.ShouldBindQuery(&queryParams); err != nil {
		logging.Logger.WithError(err).Error("Invalid query parameters")
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid query parameters",
		})
		return
	}

	filter := NewFilterFromQueryParams(queryParams)
	logging.Logger.Info("Filtering products with filter: ", filter)

	products, err := api.productService.GetProductsByFilter(filter)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to fetch products")
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to fetch products",
		})
		return
	}
	logging.Logger.Info("Products fetched successfully")

	c.JSON(http.StatusOK, products)
}

func (api *ProductAPI) CreateProduct(c *gin.Context) {
	var req messages.ProductCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logging.Logger.WithError(err).Error("Invalid request")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}
	logging.Logger.Info("Creating product with request: ", req)

	if err := api.productService.CreateProduct(req); err != nil {
		logging.Logger.WithError(err).Error("Failed to create product")
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to create product",
		})
		return
	}
	logging.Logger.Info("Product created successfully")

	c.JSON(http.StatusCreated, messages.ApiResponse{
		Code:    http.StatusCreated,
		Type:    "success",
		Message: "Product created successfully",
	})
}

func (api *ProductAPI) UpdateProduct(c *gin.Context) {
	var req messages.ProductUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logging.Logger.WithError(err).Error("Invalid request")
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}
	logging.Logger.Info("Updating product with request: ", req)

	if err := api.productService.UpdateProduct(req); err != nil {
		logging.Logger.WithError(err).Error("Failed to update product")
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to update product",
		})
		return
	}

	logging.Logger.Info("Product updated successfully")

	c.JSON(http.StatusOK, messages.ApiResponse{
		Code:    http.StatusOK,
		Type:    "success",
		Message: "Product updated successfully",
	})
}

func (api *ProductAPI) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if id <= 0 || err != nil {
		logging.Logger.WithError(err).Error("Invalid product ID")
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid product ID",
		})
		return
	}
	logging.Logger.Info("Deleting product with ID: ", id)
	if err := api.productService.DeleteProduct(int64(id)); err != nil {
		logging.Logger.WithError(err).Error("Failed to delete product")
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to delete product",
		})
		return
	}

	logging.Logger.Info("Product deleted successfully")

	c.JSON(http.StatusOK, messages.ApiResponse{
		Code:    http.StatusOK,
		Type:    "success",
		Message: "Product deleted successfully",
	})
}

func (api *ProductAPI) GetCategoryByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if id <= 0 || err != nil {
		logging.Logger.WithError(err).Error("Invalid category ID")
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid category ID",
		})
		return
	}
	logging.Logger.Info("Fetching category by ID: ", id)

	response, err := api.categoryService.GetCategoryByID(id)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to fetch category")
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to fetch category",
		})
		return
	}

	logging.Logger.Info("Category fetched successfully")
	c.JSON(http.StatusOK, response)
}

func (api *ProductAPI) GetProductsByCategory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if id <= 0 || err != nil {
		logging.Logger.WithError(err).Error("Invalid category ID")
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid category ID",
		})
		return
	}
	logging.Logger.Info("Fetching products by category ID: ", id)

	products, err := api.productService.GetProductsByFilter(messages.ProductFilter{
		CategoryIDs: []int64{id},
	})
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to fetch products")
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to fetch products",
		})
		return
	}

	logging.Logger.Info("Products fetched successfully")

	c.JSON(http.StatusOK, products)
}

func (api *ProductAPI) CreateCategory(c *gin.Context) {
	var req messages.CategoryCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logging.Logger.WithError(err).Error("Invalid request")
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}
	logging.Logger.Info("Creating category with request: ", req)

	err := api.categoryService.CreateCategory(req)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to create category")
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to create category",
		})
		return
	}

	logging.Logger.Info("Category created successfully")

	c.JSON(http.StatusCreated, messages.ApiResponse{
		Code:    http.StatusCreated,
		Type:    "success",
		Message: "Category created successfully",
	})
}

func (api *ProductAPI) UpdateCategory(c *gin.Context) {
	var req messages.CategoryUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logging.Logger.WithError(err).Error("Invalid request")
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid request",
		})
		return
	}

	logging.Logger.Info("Updating category with request: ", req)

	err := api.categoryService.UpdateCategory(req)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to update category")
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to update category",
		})
		return
	}

	logging.Logger.Info("Category updated successfully")
	c.JSON(http.StatusOK, messages.ApiResponse{
		Code:    http.StatusOK,
		Type:    "success",
		Message: "Category updated successfully",
	})
}

func (api *ProductAPI) DeleteCategory(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if id <= 0 || err != nil {
		logging.Logger.WithError(err).Error("Invalid category ID")
		c.JSON(http.StatusBadRequest, messages.ApiResponse{
			Code:    http.StatusBadRequest,
			Type:    "error",
			Message: "Invalid category ID",
		})
		return
	}

	err = api.categoryService.DeleteCategory(id)
	if err != nil {
		logging.Logger.WithError(err).Error("Failed to delete category")
		c.JSON(http.StatusInternalServerError, messages.ApiResponse{
			Code:    http.StatusInternalServerError,
			Type:    "error",
			Message: "Failed to delete category",
		})
		return
	}

	logging.Logger.Info("Category deleted successfully")
	c.JSON(http.StatusOK, messages.ApiResponse{
		Code:    http.StatusOK,
		Type:    "success",
		Message: "Category deleted successfully",
	})
}
