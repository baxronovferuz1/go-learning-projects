package repositories

import (
	"warehouse-management-system/models"

	"gorm.io/gorm"
)

type SupplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *SupplierRepository {
	return &SupplierRepository{
		db: db,
	}
}

func (r *SupplierRepository) Create(supplier *models.Supplier) error {
	return r.db.Create(supplier).Error
}

func (r *SupplierRepository) GetAll() ([]models.Supplier, error) {

	var suppliers []models.Supplier

	err := r.db.Find(&suppliers).Error

	return suppliers, err
}

func (r *SupplierRepository) GetByID(id uint) (*models.Supplier, error) {

	var supplier models.Supplier

	err := r.db.First(&supplier, id).Error
	if err != nil {
		return nil, err
	}

	return &supplier, nil
}

func (r *SupplierRepository) Update(supplier *models.Supplier) error {
	return r.db.Save(supplier).Error
}

func (r *SupplierRepository) Delete(id uint) error {
	return r.db.Delete(&models.Supplier{}, id).Error
}
