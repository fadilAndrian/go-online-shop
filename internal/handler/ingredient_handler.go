package handler

import (
	"net/http"
	"strconv"

	"github.com/fadilAndrian/go-online-shop/internal/dto"
	"github.com/fadilAndrian/go-online-shop/internal/helper"
	usecaseinterface "github.com/fadilAndrian/go-online-shop/internal/usecase/interface"
	"github.com/gin-gonic/gin"
)

type IngredientHandler struct {
	uc usecaseinterface.IngredientUsecaseInterface
}

func NewIngredientHandler(uc usecaseinterface.IngredientUsecaseInterface) *IngredientHandler {
	return &IngredientHandler{uc}
}

func (h *IngredientHandler) GetAllEquipment(c *gin.Context) {
	ingredients, err := h.uc.ListIngredient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    ingredients,
		"message": "Success",
	})
}

func (h *IngredientHandler) GetEquipmentById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	ingredient, err := h.uc.GetIngredient(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    ingredient,
		"message": "Success",
	})
}

func (h *IngredientHandler) CreateEquipment(c *gin.Context) {
	var req dto.IngredientCreateRequest

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

	ingredient, err := h.uc.CreateIngredient(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    ingredient,
		"message": "Ingredient created.",
	})
}

func (h *IngredientHandler) UpdateEquipment(c *gin.Context) {
	var req dto.IngredientUpdateRequest

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

	id, _ := strconv.Atoi(c.Param("id"))

	ingredient, err := h.uc.UpdateIngredient(int64(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    ingredient,
		"message": "Ingredient Updated",
	})
}

func (h *IngredientHandler) DeleteEquipment(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.uc.DeleteIngredient(int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ingredient Deleted",
	})
}
