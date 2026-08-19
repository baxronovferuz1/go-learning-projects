package router

import (
	"warehouse-management-system/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	categoryHandler *handlers.CategoryHandler,
	supplierHandler *handlers.SupplierHandler,
	productHandler *handlers.ProductHandler,
	stockHandler *handlers.StockHandler,
	dashboardHandler *handlers.DashboardHandler,
) *gin.Engine {

	r := gin.Default()

	// /categories endpointlari
	categories := r.Group("/categories")
	{
		// POST /categories
		categories.POST("", categoryHandler.Create)

		// GET /categories
		categories.GET("", categoryHandler.GetAll)

		// GET /categories/:id
		categories.GET("/:id", categoryHandler.GetByID)

		// PUT /categories/:id
		categories.PUT("/:id", categoryHandler.Update)

		// DELETE /categories/:id
		categories.DELETE("/:id", categoryHandler.Delete)
	}

	// =========================
	// SUPPLIER
	// =========================

	suppliers := r.Group("/suppliers")
	{
		// POST /suppliers
		suppliers.POST("", supplierHandler.Create)

		// GET /suppliers
		suppliers.GET("", supplierHandler.GetAll)

		// GET /suppliers/:id
		suppliers.GET("/:id", supplierHandler.GetByID)

		// PUT /suppliers/:id
		suppliers.PUT("/:id", supplierHandler.Update)

		// DELETE /suppliers/:id
		suppliers.DELETE("/:id", supplierHandler.Delete)
	}

	// =========================
	// PRODUCT
	// =========================

	products := r.Group("/products")
	{
		// POST /products
		products.POST("", productHandler.Create)

		// GET /products
		products.GET("", productHandler.GetAll)

		// GET /products/:id
		products.GET("/:id", productHandler.GetByID)

		// GET /products/search?keyword=iphone
		products.GET("/search", productHandler.Search)

		// GET /products/category/:category_id
		products.GET(
			"/category/:category_id",
			productHandler.FilterByCategory,
		)

		// PUT /products/:id
		products.PUT("/:id", productHandler.Update)

		// DELETE /products/:id
		products.DELETE("/:id", productHandler.Delete)
	}

	// =========================
	// STOCK
	// =========================

	stock := r.Group("/stock")
	{
		// POST /stock/movements
		stock.POST(
			"/movements",
			stockHandler.CreateMovement,
		)

		// GET /stock/history
		stock.GET(
			"/history",
			stockHandler.GetHistory,
		)
	}

	// =========================
	// DASHBOARD
	// =========================

	// GET /dashboard
	r.GET("/dashboard", dashboardHandler.GetStatistics)

	return r
}
