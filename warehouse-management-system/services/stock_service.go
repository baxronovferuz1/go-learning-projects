package services

import (
	"errors"

	"warehouse-management-system/models"
	"warehouse-management-system/repositories"
)

type StockService struct {
	repo *repositories.StockRepository
}

func NewStockService(repo *repositories.StockRepository) *StockService {
	return &StockService{
		repo: repo,
	}
}

func (s *StockService) CreateMovement(movement *models.StockMovement) error {

	if movement.ProductID == 0 {
		return errors.New("product id is required")
	}

	if movement.Quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	if movement.Type != "IN" && movement.Type != "OUT" {
		return errors.New("stock movement type must be IN or OUT")
	}

	return s.repo.CreateMovement(movement)
}

func (s *StockService) GetHistory() ([]models.StockMovement, error) {
	return s.repo.GetHistory()
}
