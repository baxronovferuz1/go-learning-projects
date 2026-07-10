package database

import (
	"library-manager/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=12345 dbname=library_manager port=5432 sllmode=disable TimeZone=Asia/Tashkent"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed", err)
	}

	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatal("auto migration failed:", err)
	}
	DB = db

	log.Println("database connection successfully")
}
