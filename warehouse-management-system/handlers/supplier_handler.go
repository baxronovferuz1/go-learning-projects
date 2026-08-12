package handlers

import (
	"net/http"
	"strconv"

	"warehouse-management-system/models"
	"warehouse-management-system/services"

	"github.com/gin-gonic/gin"
)

type SupplierHandler struct {
	service *services.SupplierService
}

func NewSupplierHandler(service *services.SupplierService) *SupplierHandler {
	return &SupplierHandler{
		service: service,
	}
}

// POST /suppliers
func (h *SupplierHandler) Create(c *gin.Context) {

	var supplier models.Supplier

	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err := h.service.Create(&supplier)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, supplier)
}

// GET /suppliers
func (h *SupplierHandler) GetAll(c *gin.Context) {

	suppliers, err := h.service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, suppliers)
}

// GET /suppliers/:id
func (h *SupplierHandler) GetByID(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid supplier id",
		})
		return
	}

	supplier, err := h.service.GetByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "supplier not found",
		})
		return
	}

	c.JSON(http.StatusOK, supplier)
}

// PUT /suppliers/:id
func (h *SupplierHandler) Update(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid supplier id",
		})
		return
	}

	var supplier models.Supplier

	if err := c.ShouldBindJSON(&supplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	supplier.ID = uint(id)

	err = h.service.Update(&supplier)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, supplier)
}

// DELETE /suppliers/:id
func (h *SupplierHandler) Delete(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid supplier id",
		})
		return
	}

	err = h.service.Delete(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "supplier not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "supplier deleted successfully",
	})
}
