package api

import "github.com/gin-gonic/gin"

type ProductAPI struct {
	// Services here.
}

func NewProductAPI() *ProductAPI {
	return &ProductAPI{}
}

func (api *ProductAPI) registerPublicOnlyRoutes(router *gin.RouterGroup) {
	// Get product here.
}

func (api *ProductAPI) registerPublicRoutes(router *gin.RouterGroup) {
	// Get all products here.
}

func (api *ProductAPI) registerPrivateRoutes(router *gin.RouterGroup) {
	// Create, update and delete product here.
}
