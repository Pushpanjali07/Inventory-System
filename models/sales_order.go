package models

import (
	"time"
)

type SalesOrder struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	ProductID  uint      `json:"product_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
	OrderDate  time.Time `json:"order_date"`
}
