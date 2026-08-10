package repositories

import (
	"warehouse-management-system/models"

	"gorm.io/gorm"
)

type DashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) *DashboardRepository {
	return &DashboardRepository{
		db: db,
	}
}

func (r *DashboardRepository) GetStatistics() (*models.DashboardResponse, error) {

	var response models.DashboardResponse

	r.db.Model(&models.Product{}).Count(&response.TotalProducts)

	r.db.Model(&models.Category{}).Count(&response.TotalCategories)

	r.db.Model(&models.Supplier{}).Count(&response.TotalSuppliers)

	r.db.Model(&models.Product{}).
		Where("stock <= ?", 5).
		Count(&response.LowStockCount)

	return &response, nil
}
