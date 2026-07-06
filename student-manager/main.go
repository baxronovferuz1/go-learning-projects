package main

import (
	"student-manager/database"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Connect()
	router := gin.Default()

	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello student manager",
	// 	})
	// })
	router.Run(":8080")

}
