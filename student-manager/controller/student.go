package controller

import (
	"net/http"
	"student-manager/database"
	"student-manager/models"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Home Page",
	})
}

// func GetStudents(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "Students List",
// 	})
// }

func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.DB.Create(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create student",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Studetn created successfully",
		"data":    student,
	})
}

func GetStudents(c *gin.Context) {
	var student []models.Student

	if err := database.DB.Find(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch student",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Student fetched succesfully",
		"data":    student,
	})
}
