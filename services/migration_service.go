package services

import (
	"cr_task1/models"
	"cr_task1/repository"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// MigrationService handles the migration of inventory from Shop-A to Shop-B.
// Single Responsibility Principle (SRP) - This service handles only the migration logic.
type MigrationService struct {
	shopBRepository repository.ProductRepository
}

// NewMigrationService creates a new MigrationService.
func NewMigrationService(shopBRepo repository.ProductRepository) *MigrationService {
	return &MigrationService{
		shopBRepository: shopBRepo,
	}
}

// MigrateInventoryFromAPI fetches inventory from Shop-A API and inserts it into Shop-B inventory.
// Adapter Pattern - This service adapts Shop-A's inventory format to Shop-B's data structure.
func (m *MigrationService) MigrateInventoryFromAPI(shopAURL string) error {
	resp, err := http.Get(shopAURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var products []models.Product
	err = json.Unmarshal(body, &products)
	if err != nil {
		return err
	}

	for _, product := range products {
		err := m.shopBRepository.CreateProduct(product)
		if err != nil {
			log.Printf("Error migrating product %d: %v\n", product.ProductID, err)
		}
	}

	return nil
}
