package router

import (
	"github.com/fadilAndrian/go-online-shop/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.RouterGroup, handler *handler.ProductHandler) {
	r.GET("/", handler.GetAll)
	r.GET("/:id", handler.Get)
	r.POST("/", handler.Create)
	r.PUT("/:id", handler.Update)
	r.DELETE("/:id", handler.Delete)
}
