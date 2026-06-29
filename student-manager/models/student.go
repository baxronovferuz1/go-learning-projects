package models

type Student struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	FullName     string `json:"full_name"`
	Phone_number int    `json:"phone_number"`
	Email        string `json:"email"`
}
