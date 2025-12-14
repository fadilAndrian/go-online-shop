package router

import (
	"github.com/fadilAndrian/go-online-shop/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterEquipmentRoutes(r *gin.RouterGroup, h *handler.EquipmentHandler) {
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
	r.POST("/", h.Create)
	r.PUT("/:id", h.Update)
	r.DELETE("/:id", h.Delete)
}
