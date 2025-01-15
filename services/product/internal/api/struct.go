package api

import (
	"math"
	"product/internal/messages"
	"strconv"
	"strings"
)

type ProductQueryParams struct {
	Category string             `form:"category"`
	Company  string             `form:"company"`
	PageSize int                `form:"pagesize"`
	Offset   int                `form:"offset"`
	MinPrice float64            `form:"minprice"`
	MaxPrice float64            `form:"maxprice"`
	Sort     string             `form:"sort"`
	SortBy   messages.SortField `form:"sortby"`
	Keyword  string             `form:"search"`
}

func NewFilterFromQueryParams(params ProductQueryParams) messages.ProductFilter {
	// Set default values if not provided
	setDefaultFilterValues(&params)

	categories := parseStringArray(params.Category)
	companyIDs := parseStringArray(params.Company)

	return messages.ProductFilter{
		CategoryIDs: categories,
		CompanyIDs:  companyIDs,
		PageSize:    params.PageSize,
		PageNumber:  params.Offset,
		MinPrice:    params.MinPrice,
		MaxPrice:    params.MaxPrice,
		Sort:        params.Sort,
		SortField:   params.SortBy,
		Keyword:     params.Keyword,
	}
}

func setDefaultFilterValues(params *ProductQueryParams) {
	if params.PageSize == 0 {
		params.PageSize = 30
	}
	if params.MinPrice > params.MaxPrice {
		params.MinPrice = 0
		params.MaxPrice = math.MaxFloat64
	}
	if params.MaxPrice == 0 {
		params.MaxPrice = math.MaxFloat64
	}
	if params.Sort != "asc" && params.Sort != "desc" {
		params.Sort = "asc"
	}
	if params.SortBy == "" {
		params.SortBy = "price"
	}
	if params.Offset < 0 {
		params.Offset = 0
	}
}

func parseStringArray(s string) []int64 {
	var result []int64
	if s == "" {
		return result
	}
	for _, id := range strings.Split(s, ",") {
		if i, err := strconv.ParseInt(id, 10, 64); err == nil {
			result = append(result, i)
		}
	}
	return result
}
