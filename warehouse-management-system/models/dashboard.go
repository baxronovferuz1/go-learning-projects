package models

type DashboardResponse struct {
	TotalProducts   int64 `json:"total_products"`
	TotalCategories int64 `json:"total_categories"`
	TotalSuppliers  int64 `json:"total_suppliers"`
	LowStockCount   int64 `json:"low_stock_count"`
}
