package controllers

import (
	"inventory-system/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateSalesOrder(c *gin.Context) {
	var order models.SalesOrder
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order.OrderDate = time.Now()

	var product models.Product
	if err := models.DB.First(&product, order.ProductID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found"})
		return
	}

	if product.Quantity < order.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
		return
	}

	order.TotalPrice = float64(order.Quantity) * product.Price
	product.Quantity -= order.Quantity
	models.DB.Save(&product)
	models.DB.Create(&order)
	c.JSON(http.StatusOK, order)
}

func GetSalesOrders(c *gin.Context) {
	var orders []models.SalesOrder
	models.DB.Find(&orders)
	c.JSON(http.StatusOK, orders)
}

func UpdateSalesOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.SalesOrder
	if err := models.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.ShouldBindJSON(&order)
	models.DB.Save(&order)
	c.JSON(http.StatusOK, order)
}

func DeleteSalesOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.SalesOrder
	result := models.DB.Delete(&order, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}
