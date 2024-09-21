package repository

import (
	"cr_task1/models"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TestCreateProduct tests the CreateProduct method in the ProductRepository.
func TestCreateProduct(t *testing.T) {
	// Create an in-memory SQLite database for testing.
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Create the products table for testing.
	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		t.Fatalf("Failed to migrate database schema: %v", err)
	}

	// Initialize the repository with the test database.
	repo := GORMProductRepository{DB: db}

	// Create a test product.
	testProduct := models.Product{
		ProductID:       1,
		ProductTitle:    "Laptop",
		ProductQuantity: 10,
		ProductPrice:    999.99,
	}

	// Call the CreateProduct method.
	err = repo.CreateProduct(testProduct)
	if err != nil {
		t.Fatalf("Failed to create product: %v", err)
	}

	// Verify that the product was inserted into the database.
	var product models.Product
	err = db.First(&product, testProduct.ProductID).Error
	if err != nil {
		t.Fatalf("Failed to retrieve product from database: %v", err)
	}

	// Check that the inserted product matches the expected values.
	if product.ProductTitle != testProduct.ProductTitle {
		t.Errorf("Expected ProductTitle to be %v, got %v", testProduct.ProductTitle, product.ProductTitle)
	}

	if product.ProductQuantity != testProduct.ProductQuantity {
		t.Errorf("Expected ProductQuantity to be %v, got %v", testProduct.ProductQuantity, product.ProductQuantity)
	}

	if product.ProductPrice != testProduct.ProductPrice {
		t.Errorf("Expected ProductPrice to be %v, got %v", testProduct.ProductPrice, product.ProductPrice)
	}
}
