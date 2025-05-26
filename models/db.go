package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // global variable to access db connection anywhere

func ConnectDatabase() {
	dsn := "root:password@tcp(127.0.0.1:3306)/inventory_db?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //Open a connection to the database using GORM
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	database.AutoMigrate(&Product{}, &SalesOrder{}, &PurchaseOrder{})
	DB = database //Save the DB reference to global variable so we can use it elsewhere
}
