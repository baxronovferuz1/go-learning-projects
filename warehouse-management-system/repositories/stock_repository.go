package repositories

import (
	"warehouse-management-system/models"

	"gorm.io/gorm"
)

type StockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) *StockRepository {
	return &StockRepository{
		db: db,
	}
}

func (r *StockRepository) CreateMovement(movement *models.StockMovement) error {
	return r.db.Create(movement).Error
}

func (r *StockRepository) GetHistory() ([]models.StockMovement, error) {

	var history []models.StockMovement

	err := r.db.
		Preload("Product").
		Find(&history).Error

	return history, err
}
