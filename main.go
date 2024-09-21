package main

import (
	"cr_task1/api"
	"cr_task1/models"
	"cr_task1/repository"
	"cr_task1/services"
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Initialize GORM with SQLite for Shop-B.
	db, err := gorm.Open(sqlite.Open("shop_b.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	// Automatically migrate the schema for `Product`.
	// This will create the `products` table if it doesn't exist.
	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatalf("Failed to migrate the database: %v\n", err)
	}

	log.Println("Products table created or already exists.")

	// Initialize repositories and APIs.
	shopBRepo := &repository.GORMProductRepository{DB: db}
	shopBAPI := &api.ShopBAPI{Repo: shopBRepo}
	migrationService := services.NewMigrationService(shopBRepo)

	// Shop-A Inventory API.
	http.HandleFunc("/shopa/inventory", api.ShopAInventoryHandler)
	// Shop-B Inventory API
	http.HandleFunc("/shopb/inventory", shopBAPI.ListProducts)

	// API to migrate Shop-A inventory to Shop-B.
	http.HandleFunc("/shopb/migrate", shopBAPI.MigrateInventory(migrationService))

	log.Println("Server started at :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
