package handler

import (
	"net/http"
	"strconv"

	"github.com/fadilAndrian/go-online-shop/internal/helper"
	"github.com/fadilAndrian/go-online-shop/internal/usecase"
	"github.com/fadilAndrian/go-online-shop/internal/usecase/dto"
	"github.com/gin-gonic/gin"
)

type EquipmentHandler struct {
	u usecase.EquipmentUsecase
}

func NewEquipmentHandler(u usecase.EquipmentUsecase) *EquipmentHandler {
	return &EquipmentHandler{u}
}

func (h *EquipmentHandler) GetAll(c *gin.Context) {
	equipments, err := h.u.ListEquipment()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data":    equipments,
		"message": "Success",
	})
}

func (h *EquipmentHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	equipment, err := h.u.GetEquipment(int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data":    equipment,
		"message": "Success",
	})
}

func (h *EquipmentHandler) Create(c *gin.Context) {
	var req dto.CreateEquipmentRequest

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

	equipment, err := h.u.CreateEquipment(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data":    equipment,
		"message": "Equipment Created",
	})
}

func (h *EquipmentHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req dto.UpdateEquipmentRequest

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

	equipment, err := h.u.UpdateEquipment(int64(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data":    equipment,
		"message": "Equipment Updated",
	})
}

func (h *EquipmentHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := h.u.DeleteEquipment(int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data":    "",
		"message": "Equipment Deleted",
	})
}
