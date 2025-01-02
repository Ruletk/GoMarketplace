package api

import (
	"net/http"
	"strconv"
	"strings"

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
	category := c.Query("category")
	categories := []int{}
	if category != "" {
		for _, s := range strings.Split(category, ",") {
			if id, err := strconv.Atoi(s); err == nil {
				categories = append(categories, id)
			}
		}
	}

	pageSize, _ := strconv.Atoi(c.DefaultQuery("pagesize", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))
	minPrice, _ := strconv.Atoi(c.DefaultQuery("minprice", "0"))
	maxPrice, _ := strconv.Atoi(c.DefaultQuery("maxprice", "0"))
	sort := c.DefaultQuery("sort", "asc")
	keyword := c.Query("search")

	// Add logic to fetch filtered products.
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
		"pageSize":   pageSize,
		"offset":     offset,
		"minPrice":   minPrice,
		"maxPrice":   maxPrice,
		"sort":       sort,
		"keyword":    keyword,
		"message":    "GetProducts not implemented yet",
	})
}

func (api *ProductAPI) CreateProduct(c *gin.Context) {
	// Add logic to create product.
	c.JSON(http.StatusOK, gin.H{"message": "CreateProduct not implemented yet"})
}

func (api *ProductAPI) UpdateProduct(c *gin.Context) {
	// Add logic to update product.
	c.JSON(http.StatusOK, gin.H{"message": "UpdateProduct not implemented yet"})
}

func (api *ProductAPI) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	// Add logic to delete product by ID.
	c.JSON(http.StatusOK, gin.H{"id": id, "message": "DeleteProduct not implemented yet"})
}

func (api *ProductAPI) GetCategoryByID(c *gin.Context) {
	id := c.Param("id")
	// Add logic to fetch category by ID.
	c.JSON(http.StatusOK, gin.H{"id": id, "message": "GetCategoryByID not implemented yet"})
}

func (api *ProductAPI) GetProductsByCategory(c *gin.Context) {
	id := c.Param("id")
	// Add logic to fetch products by category ID.
	c.JSON(http.StatusOK, gin.H{"categoryID": id, "message": "GetProductsByCategory not implemented yet"})
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
