package messages

type ProductCreateRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	CategoryID  int64   `json:"category_id" binding:"required"`
	CompanyID   int64   `json:"company_id" binding:"required"`
	DiscountID  int64   `json:"discount_id" binding:"required"`
	Quantity    int64   `json:"quantity" binding:"required"`
}

type ProductUpdateRequest struct {
	ID          int64   `json:"id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	CategoryID  int64   `json:"category_id" binding:"required"`
	DiscountID  int64   `json:"discount_id" binding:"required"`
}

type ProductResponse struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	CategoryID  int64   `json:"category_id" binding:"required"`
	CompanyID   int64   `json:"company_id" binding:"required"`
	InventoryID int64   `json:"inventory_id" binding:"required"`
	DiscountID  int64   `json:"discount_id" binding:"required"`
}

type ProductListResponse struct {
	Products []ProductResponse `json:"products"`
}

type ProductFilter struct {
	CategoryIDs []int64 `json:"category_ids"`
	CompanyIDs  []int64 `json:"company_ids"`
	MinPrice    float64 `json:"min_price"`
	MaxPrice    float64 `json:"max_price"`
	PageSize    int     `json:"page_size"`
	PageNumber  int     `json:"page_number"`
	Keyword     string  `json:"keywords"`
	Sort        string  `json:"sort_by"` // asc, desc
}

//----------------------------------------------

type CategoryCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ParentID    int64  `json:"parent_id" binding:"required"`
}

type CategoryUpdateRequest struct {
	ID          int64  `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ParentID    int64  `json:"parent_id" binding:"required"`
}

type CategoryResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	ParentID    int64  `json:"parent_id" binding:"required"`
}

type CategoryListResponse struct {
	Categories []CategoryResponse `json:"categories"`
}

//----------------------------------------------

type InventoryResponse struct {
	ID       int64 `json:"id"`
	Quantity int64 `json:"quantity" binding:"required"`
}

// ApiResponse represents a generic API response
type ApiResponse struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}
