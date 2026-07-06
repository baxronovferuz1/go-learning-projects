package database

import (
	"log"
	"student-manager/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres, password=12345 dbname=student_manager port=5432 sslmode-disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)

	}
	DB = db

	err = DB.AutoMigrate(&models.Student{})

	if err != nil {
		log.Fatal(err)
	}
}
