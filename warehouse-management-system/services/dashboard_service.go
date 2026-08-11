package services

import (
	"warehouse-management-system/models"
	"warehouse-management-system/repositories"
)

type DashboardService struct {
	repo *repositories.DashboardRepository
}

func NewDashboardService(repo *repositories.DashboardRepository) *DashboardService {
	return &DashboardService{
		repo: repo,
	}
}

func (s *DashboardService) GetStatistics() (*models.DashboardResponse, error) {

	return s.repo.GetStatistics()
}
