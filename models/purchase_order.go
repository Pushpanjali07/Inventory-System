package models

import (
	"time"
)

type PurchaseOrder struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Supplier  string    `json:"supplier"`
	OrderDate time.Time `json:"order_date"`
}
