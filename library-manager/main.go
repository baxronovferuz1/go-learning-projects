package main

import (
	"library-manager/database"
	"library-manager/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectDB()

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")
}
