// main()

// --LoadConfig() -config
// --Connect()    database/postgres
// --Migrate()
// --Gin Router
// --Run()

package main

import (
	"api-aggregator/config"
	"api-aggregator/database"
	"api-aggregator/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect(conf)
	if err != nil {
		log.Fatal(err)
	}

	if err := database.Migrate(db); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	router.SetupRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
