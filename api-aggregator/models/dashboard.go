package models

import (
	"api-aggregator/repositories"
)

type DashboardResponse struct {
	userRepo          *repositories.UserRepository
	orderRepo         *repositories.OrderRepository
	notificationsRepo *repositories.NotificationRepository
}
