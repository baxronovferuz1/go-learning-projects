package services

import (
	"errors"
	"warehouse-management-system/models"
	"warehouse-management-system/repositories"
)

type SupplierService struct {
	repo *repositories.SupplierRepository
}

func NewSupplierService(repo *repositories.SupplierRepository) *SupplierService {
	return &SupplierService{
		repo: repo,
	}
}

func (s *SupplierService) Create(supplier *models.Supplier) error {

	if supplier.Name == "" {
		return errors.New("supplier name is required")
	}

	if supplier.Phone == "" {
		return errors.New("phone is required")
	}

	return s.repo.Create(supplier)
}

func (s *SupplierService) GetAll() ([]models.Supplier, error) {
	return s.repo.GetAll()
}

func (s *SupplierService) GetByID(id uint) (*models.Supplier, error) {
	return s.repo.GetByID(id)
}

func (s *SupplierService) Update(supplier *models.Supplier) error {

	if supplier.Name == "" {
		return errors.New("supplier name is required")
	}

	if supplier.Phone == "" {
		return errors.New("phone is required")
	}

	return s.repo.Update(supplier)
}

func (s *SupplierService) Delete(id uint) error {
	return s.repo.Delete(id)
}
