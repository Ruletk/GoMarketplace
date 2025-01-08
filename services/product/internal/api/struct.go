package api

import (
	"product/internal/messages"
	"strconv"
	"strings"
)

type ProductQueryParams struct {
	Category string  `form:"category"`
	Company  string  `form:"company"`
	PageSize int     `form:"pagesize"`
	Offset   int     `form:"offset"`
	MinPrice float64 `form:"minprice"`
	MaxPrice float64 `form:"maxprice"`
	Sort     string  `form:"sort"`
	Keyword  string  `form:"search"`
}

func NewFilterFromQueryParams(params ProductQueryParams) messages.ProductFilter {
	// Set default values if not provided
	if params.PageSize == 0 {
		params.PageSize = 30
	}
	if params.MaxPrice == 0 {
		params.MaxPrice = 999999999999999
	}
	if params.Sort == "" {
		params.Sort = "asc"
	}

	var categories []int64
	if params.Category != "" {
		for _, s := range strings.Split(params.Category, ",") {
			if id, err := strconv.ParseInt(s, 10, 64); err == nil {
				categories = append(categories, id)
			}
		}
	}

	var companyIDs []int64
	if params.Company != "" {
		for _, s := range strings.Split(params.Company, ",") {
			if id, err := strconv.ParseInt(s, 10, 64); err == nil {
				companyIDs = append(companyIDs, id)
			}
		}
	}

	return messages.ProductFilter{
		CategoryIDs: categories,
		PageSize:    params.PageSize,
		PageNumber:  params.Offset,
		MinPrice:    params.MinPrice,
		MaxPrice:    params.MaxPrice,
		Sort:        params.Sort,
		Keyword:     params.Keyword,
	}
}
