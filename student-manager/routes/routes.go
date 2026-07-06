package routes

import (
	"student-manager/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouters(router *gin.Engine) {
	// router.GET("/", controller.Home)
	router.POST("/students", controller.CreateStudent)
	router.GET("/students", controller.GetStudents)
	router.GET("/students/:id", controller.GetStudentByID)
	router.PUT("students/:id", controller.UpdateStudent)
	router.DELETE("/student/:id", controller.DeletedStudent)
}
