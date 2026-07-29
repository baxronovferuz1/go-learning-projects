package models

import "time"

type StockMovement struct {
	ID uint `gorm:"primaryKey" json:"id"`

	ProductID uint    `gorm:"not null" json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`

	Type      string    `gorm:"size:10;not null" json:"type"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	Note      string    `gorm:"size:255" json:"note"`
	CreatedAt time.Time `json:"created_at"`
}
