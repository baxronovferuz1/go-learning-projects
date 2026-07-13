package models

import "time"

type Order struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Product   string    `json:"product" gorm:"type:varchar(100);not null"`
	Price     float64   `json:"price" gorm:"type:numeric(12,2);not null"`
	Status    string    `json:"status" gorm:"type:varchar(40);default:'peding'"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
