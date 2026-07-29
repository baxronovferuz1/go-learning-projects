package models

type DashboardResponse struct {
	TotalProducts   int `json:"total_products"`
	ToTalCategories int `json:"total_categories"`
	TotalSuppliers  int `json:"total_suppliers"`
	LowStockCount   int `json:"low_stock_count"`
}
