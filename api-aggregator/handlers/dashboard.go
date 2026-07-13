package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDashboard(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Dashboad Api",
	})
}
