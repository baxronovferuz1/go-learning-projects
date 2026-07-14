// main()

// --LoadConfig() -config
// --Connect()    database/postgres
// --Migrate()
// --Gin Router
// --Run()

package main

import (
	"api-aggregator/router"

	"github.com/gin-gonic/gin"
)

func main() {

	// database.ConnectDB()
	r := gin.Default()

	router.SetupRoutes(r)

	r.Run(":8080")

}
