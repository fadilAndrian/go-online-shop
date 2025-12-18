package router

import (
	"github.com/fadilAndrian/go-online-shop/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterIngredientRoutes(r *gin.RouterGroup, h *handler.IngredientHandler) {
	r.GET("/", h.GetAllEquipment)
	r.GET("/:id", h.GetEquipmentById)
	r.POST("/", h.CreateEquipment)
	r.PUT("/:id", h.UpdateEquipment)
	r.DELETE("/:id", h.DeleteEquipment)
}
