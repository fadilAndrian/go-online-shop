package handler

import (
	"net/http"
	"strconv"

	"github.com/fadilAndrian/go-online-shop/internal/dto"
	"github.com/fadilAndrian/go-online-shop/internal/helper"
	"github.com/fadilAndrian/go-online-shop/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	uc *usecase.ProductUsecase
}

func NewProductHandler(uc *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{uc}
}

func (handler *ProductHandler) GetAll(c *gin.Context) {
	products, err := handler.uc.ListProduct()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Success",
		"data":    products,
	})
}

func (handler *ProductHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := handler.uc.GetProduct(int64(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Success",
		"data":    product,
	})
}

func (handler *ProductHandler) Create(c *gin.Context) {
	var req dto.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	if err := helper.ValidateRequest(req, c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	product, err := handler.uc.CreateProduct(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Product Created.",
		"data":    product,
	})
}

func (handler *ProductHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req dto.UpdateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	if err := helper.ValidateRequest(req, c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	result, err := handler.uc.UpdateProduct(int64(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Product Updated.",
		"data":    result,
	})
}

func (handler *ProductHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := handler.uc.DeleteProduct(int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "Product Deleted.",
	})
}
