package models

import (
	"time"
)

type PurchaseOrder struct {
	ProductID uint      `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Supplier  string    `json:"supplier"`
	OrderDate time.Time `json:"order_date"`
}
