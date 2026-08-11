package services

import (
	"errors"

	"warehouse-management-system/models"
	"warehouse-management-system/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) Create(product *models.Product) error {

	if product.Name == "" {
		return errors.New("product name is required")
	}

	if product.Price <= 0 {
		return errors.New("product price must be greater than zero")
	}

	if product.Stock < 0 {
		return errors.New("stock cannot be negative")
	}

	return s.repo.Create(product)
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *ProductService) GetByID(id uint) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Search(keyword string) ([]models.Product, error) {

	if keyword == "" {
		return s.repo.GetAll()
	}

	return s.repo.Search(keyword)
}

func (s *ProductService) FilterByCategory(categoryID uint) ([]models.Product, error) {

	if categoryID == 0 {
		return nil, errors.New("category id is required")
	}

	return s.repo.FilterByCategory(categoryID)
}

func (s *ProductService) Update(product *models.Product) error {

	if product.Name == "" {
		return errors.New("product name is required")
	}

	if product.Price <= 0 {
		return errors.New("product price must be greater than zero")
	}

	if product.Stock < 0 {
		return errors.New("stock cannot be negative")
	}

	return s.repo.Update(product)
}

func (s *ProductService) Delete(id uint) error {

	if id == 0 {
		return errors.New("product id is required")
	}

	return s.repo.Delete(id)
}
