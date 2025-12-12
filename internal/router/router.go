package router

import (
	"github.com/fadilAndrian/go-online-shop/internal/handler"
	"github.com/gin-gonic/gin"
)

func Routes(
	ProductHandler *handler.ProductHandler,
) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	ProductGroup := api.Group("/products")
	RegisterProductRoutes(ProductGroup, ProductHandler)

	return r
}
