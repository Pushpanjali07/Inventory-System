package controllers

import (
	"inventory-system/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetInventory(c *gin.Context) {
	var products []models.Product
	models.DB.Select("id", "name", "quantity").Find(&products)
	c.JSON(http.StatusOK, products)
}
