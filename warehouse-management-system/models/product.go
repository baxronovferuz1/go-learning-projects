package models

import (
	"time"
)

type Product struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `gorm:"size:100;not null" json:"name"`
	Description string  `gorm:"size:500" json:"description"`
	Price       float64 `gorm:"not null" json:"price"`
	Stock       int     `gorm:"default:0" json:"stock"`

	CategoryID uint     `json:"category_id" gorm:"not null"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryID"`
	SupplierID uint     `json:"supplier_id" gorm:"not null"`
	Supplier   Supplier `json:"supplier" gorm:"foreignKey:SupplierID"`

	StockMovements []StockMovement `json:"stock_movements" gorm:"foreignKey:ProductID"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
