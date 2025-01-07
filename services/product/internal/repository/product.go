package repository

import (
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
}

func (Product) TableName() string {
	return "products"
}

type ProductRepository interface {
	GetByID(id int64) (*Product, error)
	GetByCompanyID(id int64) ([]*Product, error)
	GetByName(name string) ([]*Product, error)
	GetByPage(pageSize int, offset int) ([]*Product, error)
	GetByFilter(filter *messages.ProductFilter) ([]*Product, error)
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
	return &productRepository{
		db: db,
	}
}
func (p productRepository) GetByID(id int64) (*Product, error) {
	var product Product
	err := p.db.Where("product_id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p productRepository) GetByCompanyID(id int64) ([]*Product, error) {
	var products []*Product
	err := p.db.Where("company_id = ?", id).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p productRepository) GetByName(name string) ([]*Product, error) {
	var products []*Product
	err := p.db.Where("name = ?", name).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p productRepository) GetByPage(pageSize int, offset int) ([]*Product, error) {
	var products []*Product
	err := p.db.Limit(pageSize).Offset(offset).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p productRepository) GetByFilter(filter *messages.ProductFilter) ([]*Product, error) {

	var products []*Product

	query := p.db.Limit(filter.PageSize).Offset(filter.PageNumber)
	query = query.Where("price >= ?", filter.MinPrice)
	query = query.Where("price <= ?", filter.MaxPrice)

	if filter.Keyword != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Keyword+"%")
	}
	if filter.Sort == "asc" {
		query = query.Order("price asc")
	} else {
		query = query.Order("price desc")
	}
	if len(filter.CompanyIDs) > 0 {
		query = query.Where("company_id IN ?", filter.CompanyIDs)
	}
	if len(filter.CategoryIDs) > 0 {
		query = query.Where("category_id IN ?", filter.CategoryIDs)
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p productRepository) Create(product *Product) error {
	return p.db.Create(product).Error
}

func (p productRepository) Update(product *Product) error {
	return p.db.Model(&Product{}).Where("product_id = ?", product.ProductID).Updates(product).Error
}

func (p productRepository) DeleteByID(id int64) error {
	return p.db.Where("product_id = ?", id).Delete(&Product{}).Error
}

func (p productRepository) DeleteAllByName(name string) error {
	return p.db.Where("name = ?", name).Delete(&Product{}).Error
}

func (p productRepository) DeleteAllByCategoryID(id int64) error {
	return p.db.Where("category_id = ?", id).Delete(&Product{}).Error
}

func (p productRepository) DeleteAllByCompanyID(id int64) error {
	return p.db.Where("company_id = ?", id).Delete(&Product{}).Error
}
