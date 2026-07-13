package router

import (
	"api-aggregator/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/dashboard", handlers.GetDashboard)
}
