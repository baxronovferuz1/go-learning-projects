package handlers

import (
	"net/http"
	"strconv"

	"warehouse-management-system/models"
	"warehouse-management-system/services"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

// POST /categories

func (h *CategoryHandler) Create(c *gin.Context) {

	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err := h.service.Create(&category)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, category)

}

// GET /categories

func (h *CategoryHandler) GetAll(c *gin.Context) {

	categories, err := h.service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// GET /categories/:id

func (h *CategoryHandler) GetByID(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid category id",
		})
		return
	}

	category, err := h.service.GetByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "category not found",
		})
		return
	}

	c.JSON(http.StatusOK, category)
}

// PUT /categories/:id

func (h *CategoryHandler) Update(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid category id",
		})
		return
	}

	var category models.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	category.ID = uint(id)

	err = h.service.Update(&category)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, category)
}

// DELETE /categories/:id
func (h *CategoryHandler) Delete(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid category id",
		})
		return
	}

	err = h.service.Delete(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "category not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "category deleted successfully",
	})
}
