package routes

import (
	"library-manager/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	bookRoutes := router.Group("/books")
	{
		bookRoutes.POST("", controller.CreateBook)
		bookRoutes.GET("", controller.GetBooks)
		bookRoutes.GET("/:id", controller.GetBooks)
		bookRoutes.PUT("/:id", controller.UpdateBook)
		bookRoutes.DELETE("/:id", controller.DeleteBook)

	}
}
