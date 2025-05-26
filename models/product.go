package models

// We import GORM to use its built-in model fields and ORM features

type Product struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}
