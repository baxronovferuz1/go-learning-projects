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
	var students []models.Student

	if err := database.DB.Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch student",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Student fetched succesfully",
		"data":    students,
	})
}

func GetStudentByID(c *gin.Context) {
	id := c.Param("id")

	var student models.Student

	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully found",
		"data":    student,
	})
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")

	var student models.Student

	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})
		return
	}

	var input models.Student
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	student.FullName = input.FullName
	student.Phone_number = input.Phone_number
	student.Email = input.Email

	if err := database.DB.Save(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to upadte student",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "student uptated successfully",
		"data":    student,
	})
}

func DeletedStudent(c *gin.Context) {
	id := c.Param("id")

	var student models.Student

	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})
		return
	}

	if err := database.DB.Delete(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete student",
		})
		return

	}
	c.JSON((http.StatusOK), gin.H{
		"message": "Student successfully deleted",
	})
}
