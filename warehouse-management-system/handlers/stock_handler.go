package handlers

import (
	"net/http"

	"warehouse-management-system/models"
	"warehouse-management-system/services"

	"github.com/gin-gonic/gin"
)

type StockHandler struct {
	service *services.StockService
}

func NewStockHandler(service *services.StockService) *StockHandler {
	return &StockHandler{
		service: service,
	}
}

// POST /stock/movements
func (h *StockHandler) CreateMovement(c *gin.Context) {

	var movement models.StockMovement

	if err := c.ShouldBindJSON(&movement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err := h.service.CreateMovement(&movement)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, movement)
}

// GET /stock/history
func (h *StockHandler) GetHistory(c *gin.Context) {

	history, err := h.service.GetHistory()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, history)
}
