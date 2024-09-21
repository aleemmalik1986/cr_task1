package repository

import (
	"cr_task1/models"

	"gorm.io/gorm"
)

// ProductRepository defines the CRUD operations for managing Shop-B's inventory.
// Interface Segregation Principle (ISP) - The repository interface is specific and provides only the methods needed for Shop-B.
type ProductRepository interface {
	CreateProduct(product models.Product) error
	ListProducts() ([]models.Product, error)
}

// GORMProductRepository is the GORM-based implementation of ProductRepository.
// Repository Pattern - This pattern abstracts the data access logic and separates it from the business logic.
type GORMProductRepository struct {
	DB *gorm.DB
}

// CreateProduct adds a new product to Shop-B's inventory.
func (r *GORMProductRepository) CreateProduct(product models.Product) error {
	return r.DB.Create(&product).Error
}

// ListProducts returns all products in Shop-B's inventory.
func (r *GORMProductRepository) ListProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Find(&products).Error
	return products, err
}
