package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"cr_task1/models"
)

// ShopAAPI provides methods to interact with Shop-A's inventory.
// Single Responsibility Principle (SRP) - This struct handles fetching Shop-A's inventory.
type ShopAAPI struct {
	FilePath string
}

// GetInventory reads the inventory from a JSON file.
// Adapter Pattern - This adapts Shop-A's inventory format (JSON) to a structure usable by Shop-B's migration system.
func (s *ShopAAPI) GetInventory() ([]models.Product, error) {
	jsonFile, err := os.Open(s.FilePath)
	if err != nil {
		log.Printf("Error opening file: %v\n", err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Printf("Error reading file: %v\n", err)
		return nil, err
	}

	var products []models.Product
	err = json.Unmarshal(byteValue, &products)
	if err != nil {
		log.Printf("Error unmarshaling JSON: %v\n", err)
		return nil, err
	}

	return products, nil
}

// ShopAInventoryHandler serves as the HTTP handler for Shop-A's API.
func ShopAInventoryHandler(w http.ResponseWriter, r *http.Request) {
	shopA := ShopAAPI{FilePath: "data/shop_A_inventory.json"}
	inventory, err := shopA.GetInventory()
	if err != nil {
		http.Error(w, "Failed to retrieve inventory", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(inventory)
}
