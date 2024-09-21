package api

import (
	"cr_task1/repository"
	"cr_task1/services"
	"encoding/json"
	"net/http"
)

// ShopBAPI exposes the CRUD operations for Shop-B's inventory.
// Single Responsibility Principle (SRP) - This API handles only HTTP requests and responses related to CRUD operations.
type ShopBAPI struct {
	Repo repository.ProductRepository
}

// ListProducts retrieves all products in Shop-B's inventory.
func (api *ShopBAPI) ListProducts(w http.ResponseWriter, r *http.Request) {
	products, err := api.Repo.ListProducts()
	if err != nil {
		http.Error(w, "Unable to fetch products no product exist in ShopB ", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

// MigrateInventory handles triggering the migration from Shop-A's API.
func (api *ShopBAPI) MigrateInventory(service *services.MigrationService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shopAURL := r.URL.Query().Get("shopAURL")
		err := service.MigrateInventoryFromAPI(shopAURL)
		if err != nil {
			http.Error(w, "Failed to migrate inventory", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode("Inventory migrated successfully")
	}
}
