package routes

import (
	"student-manager/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouters(router *gin.Engine) {
	router.GET("/", controller.Home)
	router.GET("/students", controller.GetStudents)
}
