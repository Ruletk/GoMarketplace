package repository

import (
	"github.com/Ruletk/GoMarketplace/pkg/logging"
	"gorm.io/gorm"
	"product/internal/messages"
)

type Product struct {
	ProductID   int64   `json:"product_id" gorm:"primaryKey" gorm:"column:product_id"`
	CompanyID   int64   `json:"company_id" gorm:"column:company_id"`
	Name        string  `json:"name" gorm:"column:name"`
	Description string  `json:"description" gorm:"column:description"`
	Price       float64 `json:"price" gorm:"column:price"`
	CategoryID  int64   `json:"category_id" gorm:"column:category_id"`
	InventoryID int64   `json:"inventory_id" gorm:"column:inventory_id"`
	DiscountID  int64   `json:"discount_id" gorm:"column:discount_id"`
	CreatedAt   string  `json:"created_at" gorm:"column:created_at" gorm:"autoCreateTime"`
	UpdatedAt   string  `json:"updated_at" gorm:"column:updated_at" gorm:"autoUpdateTime"`
}

func (Product) TableName() string {
	return "products"
}

type ProductRepository interface {
	GetByID(id int64) (*Product, error)
	GetByCompanyID(id int64) ([]*Product, error)
	GetByName(name string) ([]*Product, error)
	GetByPage(pageSize int, offset int) ([]*Product, error)
	GetByFilter(filter *messages.ProductFilter) (products []*Product, count int64, err error)
	Create(product *Product) error
	Update(product *Product) error
	DeleteByID(id int64) error
	DeleteAllByName(name string) error
	DeleteAllByCategoryID(id int64) error
	DeleteAllByCompanyID(id int64) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	logging.Logger.Info("Creating new product repository")
	return &productRepository{
		db: db,
	}
}
func (p productRepository) GetByID(id int64) (*Product, error) {
	logging.Logger.Debug("Getting product by id: ", id)
	var product Product
	err := p.db.Where("product_id = ?", id).First(&product).Error
	if err != nil {
		logging.Logger.WithError(err).Error("Error getting product by id: ", id)
		return nil, err
	}
	logging.Logger.Info("Product found: ", product)
	return &product, nil
}

func (p productRepository) GetByCompanyID(id int64) ([]*Product, error) {
	logging.Logger.Debug("Getting products by company id: ", id)
	var products []*Product
	err := p.db.Where("company_id = ?", id).Find(&products).Error
	if err != nil {
		logging.Logger.WithError(err).Error("Error getting products by company id: ", id)
		return nil, err
	}
	logging.Logger.Info("Products found: ", products)
	return products, nil
}

func (p productRepository) GetByName(name string) ([]*Product, error) {
	logging.Logger.Debug("Getting products by name: ", name)
	var products []*Product
	err := p.db.Where("name = ?", name).Find(&products).Error
	if err != nil {
		logging.Logger.WithError(err).Error("Error getting products by name: ", name)
		return nil, err
	}
	logging.Logger.Info("Products found: ", products)
	return products, nil
}

func (p productRepository) GetByPage(pageSize int, offset int) ([]*Product, error) {
	logging.Logger.Debug("Getting products by page")
	var products []*Product
	err := p.db.Limit(pageSize).Offset(offset).Find(&products).Error
	if err != nil {
		logging.Logger.WithError(err).Error("Error getting products by page")
		return nil, err
	}
	logging.Logger.Info("Products found: ", products)
	return products, nil
}

func (p productRepository) GetByFilter(filter *messages.ProductFilter) (products []*Product, count int64, err error) {
	logging.Logger.Debug("Getting productArray by filter")
	var productArray []*Product
	var totalCount int64

	query := p.db.Model(&Product{})
	query = query.Where("price >= ?", filter.MinPrice)
	query = query.Where("price <= ?", filter.MaxPrice)

	var sortField messages.SortField
	switch filter.SortField {
	case messages.SortByName:
		sortField = "name"
	case messages.SortByPrice:
		sortField = "price"
	case messages.SortByPopularity:
		sortField = "price" // TODO: Implement popularity
	case messages.SortByDate:
		sortField = "created_at"
	default:
		sortField = "price"
	}

	// Already ensured that sort is either ASC or DESC. Safe to use string concatenation
	query = query.Order(string(sortField) + " " + filter.Sort)

	logging.Logger.Debug("Filtering by price: ", filter.MinPrice, filter.MaxPrice)

	if filter.Keyword != "" {
		logging.Logger.Debug("Filtering by keyword: ", filter.Keyword)
		query = query.Where("name ILIKE ?", "%"+filter.Keyword+"%")
	}

	if len(filter.CompanyIDs) > 0 {
		logging.Logger.Debug("Filtering by company ids: ", filter.CompanyIDs)
		query = query.Where("company_id IN ?", filter.CompanyIDs)
	}
	if len(filter.CategoryIDs) > 0 {
		logging.Logger.Debug("Filtering by category ids: ", filter.CategoryIDs)
		query = query.Where("category_id IN ?", filter.CategoryIDs)
	}
	query.Count(&totalCount)

	query = query.Limit(filter.PageSize).Offset(filter.PageNumber)
	logging.Logger.Debug("Executing query, total count: ", totalCount)
	if err := query.Find(&productArray).Error; err != nil {
		logging.Logger.WithError(err).Error("Error getting productArray by filter")
		return nil, 0, err
	}
	logging.Logger.Info("Products found. Total: ", len(productArray))
	return productArray, totalCount, nil
}

func (p productRepository) Create(product *Product) error {
	logging.Logger.Debug("Creating product: ", product)
	if err := p.db.Create(product).Error; err != nil {
		logging.Logger.WithError(err).Error("Error creating product: ", product)
		return err
	}
	logging.Logger.Info("Product created: ", product)
	return nil
}

func (p productRepository) Update(product *Product) error {
	logging.Logger.Debug("Updating product: ", product)
	if err := p.db.Model(&Product{}).Where("product_id = ?", product.ProductID).Updates(product).Error; err != nil {
		logging.Logger.WithError(err).Error("Error updating product: ", product)
		return err
	}
	logging.Logger.Info("Product updated: ", product)
	return nil
}

func (p productRepository) DeleteByID(id int64) error {
	if err := p.db.Where("product_id = ?", id).Delete(&Product{}).Error; err != nil {
		logging.Logger.WithError(err).Error("Error deleting product by id: ", id)
		return err
	}
	logging.Logger.Info("Product deleted by id: ", id)
	return nil
}

func (p productRepository) DeleteAllByName(name string) error {
	logging.Logger.Debug("Deleting products by name: ", name)
	if err := p.db.Where("name = ?", name).Delete(&Product{}).Error; err != nil {
		logging.Logger.WithError(err).Error("Error deleting products by name: ", name)
		return err
	}
	logging.Logger.Info("Products deleted by name: ", name)
	return nil
}

func (p productRepository) DeleteAllByCategoryID(id int64) error {
	logging.Logger.Debug("Deleting products by category id: ", id)
	if err := p.db.Where("category_id = ?", id).Delete(&Product{}).Error; err != nil {
		logging.Logger.WithError(err).Error("Error deleting products by category id: ", id)
		return err
	}
	logging.Logger.Info("Products deleted by category id: ", id)
	return nil
}

func (p productRepository) DeleteAllByCompanyID(id int64) error {
	logging.Logger.Debug("Deleting products by company id: ", id)
	if err := p.db.Where("company_id = ?", id).Delete(&Product{}).Error; err != nil {
		logging.Logger.WithError(err).Error("Error deleting products by company id: ", id)
		return err
	}
	logging.Logger.Info("Products deleted by company id: ", id)
	return nil
}
