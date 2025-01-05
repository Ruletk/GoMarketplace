package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"product/internal/api"
	"product/internal/repository"
	"product/internal/service"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token", "Cookie"},
		AllowCredentials: true,
	}))

	db := &gorm.DB{} // Just for preventing error

	// Initialize repositories
	categoryRepo := repository.NewCategoryRepository(db)
	productRepo := repository.NewProductRepository(db)
	inventoryRepo := repository.NewInventoryRepository(db)
	// Initialize services
	categoryService := service.NewCategoryService(categoryRepo)
	productService := service.NewProductService(productRepo)
	inventoryService := service.NewInventoryService(inventoryRepo)

	// Initialize APIs
	productAPI := api.NewProductAPI(inventoryService, productService, categoryService)

	// Register routes
	public := r.Group("/")
	publicOnly := r.Group("/")
	private := r.Group("/")

	productAPI.RegisterPublicRoutes(public)
	productAPI.RegisterPublicOnlyRoutes(publicOnly)
	productAPI.RegisterPrivateRoutes(private)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
