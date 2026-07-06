package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	// ID           uint   `json:"id" gorm:"primaryKey"`
	FullName     string `json:"full_name" binding:"required"`
	Phone_number string `json:"phone_number" binding:"required"`
	Email        string `json:"email" binding:"required,email" gorm:"unique"`
	Age          int    `json:"age" binding:"required,gte=1"`
}
