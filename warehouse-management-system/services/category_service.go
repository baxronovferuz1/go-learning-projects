package services

import (
	"errors"
	"warehouse-management-system/models"
	"warehouse-management-system/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (s *CategoryService) Create(category *models.Category) error {

	if category.Name == "" {
		return errors.New("category name is required")
	}

	return s.repo.Create(category)
}

func (s *CategoryService) GetAll() ([]models.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) GetByID(id uint) (*models.Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) Update(category *models.Category) error {

	if category.Name == "" {
		return errors.New("category name is required")
	}

	return s.repo.Update(category)
}

func (s *CategoryService) Delete(id uint) error {
	return s.repo.Delete(id)
}
