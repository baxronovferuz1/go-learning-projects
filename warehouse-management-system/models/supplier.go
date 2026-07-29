package models

import "time"

type Supplier struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Phone     string    `gorm:"size:20;not null;unique" json:"phone"`
	Email     string    `gorm:"size:100;unique" json:"email"`
	Address   string    `gorm:"size:255" json:"address"`
	Products  []Product `gorm:"foreignKey:SupplierID" json:"products"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
