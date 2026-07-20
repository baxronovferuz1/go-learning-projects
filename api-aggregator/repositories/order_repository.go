package repositories

import (
	"api-aggregator/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) GetOrdersByUserID(userID uint) ([]models.Order, error) {

	var orders []models.Order

	err := r.db.Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil
}
