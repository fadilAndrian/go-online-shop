package router

import (
	"github.com/fadilAndrian/go-online-shop/internal/handler"
	"github.com/gin-gonic/gin"
)

func Routes(
	ProductHandler *handler.ProductHandler,
	EquipmentHandler *handler.EquipmentHandler,
	IngredientHandler *handler.IngredientHandler,
) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	ProductGroup := api.Group("/products")
	RegisterProductRoutes(ProductGroup, ProductHandler)

	EquipmentGroup := api.Group("/equipments")
	RegisterEquipmentRoutes(EquipmentGroup, EquipmentHandler)

	IngredientGroup := api.Group("/ingredients")
	RegisterIngredientRoutes(IngredientGroup, IngredientHandler)

	return r
}
