package models

type Book struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Title         string `json:"title" gorm:"not null"`
	Author        string `json:"author" gorm:"not null"`
	Category      string `json:"category" gorm:"not null"`
	PublishedYear int    `json:"published_year" gorm:"not null"`
	Available     bool   `json:"available" gorm:"default:true"`
}
