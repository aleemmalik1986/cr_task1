package models

import "gorm.io/gorm"

// Product represents the product entity for both Shop-A and Shop-B.
// Single Responsibility Principle (SRP) - This struct is responsible for representing product data in both systems.
type Product struct {
	gorm.Model
	ProductID       int     `json:"product_id" gorm:"primaryKey"`
	ProductTitle    string  `json:"product_title"`
	ProductQuantity int     `json:"product_quantity"`
	ProductPrice    float64 `json:"product_price"`
}
