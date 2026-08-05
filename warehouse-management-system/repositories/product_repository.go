package repositories

import (
	"warehouse-management-system/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {

	var products []models.Product

	err := r.db.
		Preload("Category").
		Preload("Supplier").
		Find(&products).Error

	return products, err
}

func (r *ProductRepository) GetByID(id uint) (*models.Product, error) {

	var product models.Product

	err := r.db.
		Preload("Category").
		Preload("Supplier").
		First(&product, id).Error

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) Search(keyword string) ([]models.Product, error) {

	var products []models.Product

	err := r.db.
		Where("name ILIKE ?", "%"+keyword+"%").
		Preload("Category").
		Preload("Supplier").
		Find(&products).Error

	return products, err
}

func (r *ProductRepository) FilterByCategory(categoryID uint) ([]models.Product, error) {

	var products []models.Product

	err := r.db.
		Where("category_id = ?", categoryID).
		Preload("Category").
		Preload("Supplier").
		Find(&products).Error

	return products, err
}

func (r *ProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}
