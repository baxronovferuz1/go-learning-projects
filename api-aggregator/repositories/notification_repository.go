package repositories

import (
	"api-aggregator/models"

	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{
		db: db,
	}
}

func (r *NotificationRepository) GetNotificationsByUserID(userID uint) ([]models.Notification, error) {

	var notifications []models.Notification

	err := r.db.Where("user_id = ?", userID).Find(&notifications).Error
	if err != nil {
		return nil, err
	}

	return notifications, nil
}
