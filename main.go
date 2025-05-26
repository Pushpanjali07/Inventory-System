package main

import (
	"inventory-system/models"
	"inventory-system/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()
	routes.SetupRoutes(router)
	router.Run(":8080")
}
