package handlers

import (
	"net/http"

	"warehouse-management-system/services"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	service *services.DashboardService
}

func NewDashboardHandler(service *services.DashboardService) *DashboardHandler {
	return &DashboardHandler{
		service: service,
	}
}

// GET /dashboard
func (h *DashboardHandler) GetStatistics(c *gin.Context) {

	statistics, err := h.service.GetStatistics()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, statistics)
}
