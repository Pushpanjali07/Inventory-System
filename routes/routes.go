package routes

import (
	"inventory-system/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/products", controllers.CreateProduct)
		api.GET("/products", controllers.GetProducts)
		api.PUT("/products/:id", controllers.UpdateProduct)
		api.DELETE("/products/:id", controllers.DeleteProduct)

		// Sales Order Routes
		api.POST("/sales-orders", controllers.CreateSalesOrder)
		api.GET("/sales-orders", controllers.GetSalesOrders)
		api.GET("/sales-orders/:id", controllers.UpdateSalesOrder)
		api.DELETE("/sales-orders/:id", controllers.DeleteSalesOrder)

		// Purchase Order Routes
		api.POST("/purchase-orders", controllers.CreatePurchaseOrder)
		api.GET("/purchase-orders", controllers.GetPurchaseOrders)
		api.GET("/purchase-orders/:id", controllers.UpdatePurchaseOrder)
		api.DELETE("/purchase-orders/:id", controllers.DeletePurchaseOrder)

		// Inventory Route
		api.GET("/inventory", controllers.GetInventory)
	}
}
