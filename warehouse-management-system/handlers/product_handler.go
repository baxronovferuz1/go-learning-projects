package handlers

import (
	"net/http"
	"strconv"

	"warehouse-management-system/models"
	"warehouse-management-system/services"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

// POST /products
func (h *ProductHandler) Create(c *gin.Context) {

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	err := h.service.Create(&product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// GET /products
func (h *ProductHandler) GetAll(c *gin.Context) {

	products, err := h.service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GET /products/:id
func (h *ProductHandler) GetByID(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product id",
		})
		return
	}

	product, err := h.service.GetByID(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

// GET /products/search?keyword=iphone
func (h *ProductHandler) Search(c *gin.Context) {

	keyword := c.Query("keyword")

	products, err := h.service.Search(keyword)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GET /products/category/:category_id
func (h *ProductHandler) FilterByCategory(c *gin.Context) {

	categoryID, err := strconv.ParseUint(
		c.Param("category_id"),
		10,
		64,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid category id",
		})
		return
	}

	products, err := h.service.FilterByCategory(uint(categoryID))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

// PUT /products/:id
func (h *ProductHandler) Update(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product id",
		})
		return
	}

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	product.ID = uint(id)

	err = h.service.Update(&product)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DELETE /products/:id
func (h *ProductHandler) Delete(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product id",
		})
		return
	}

	err = h.service.Delete(uint(id))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product deleted successfully",
	})
}
