# Shop Migration Project

## Problem Statement
Shop-A is selling all of it's inventory and Shop-B is buying all inventory from Shop A. Shop-A has different backend to store its inventory and Shop-B has different. Write a software where Shop-B can migrate all inventory from Shop-A.

### Functional Requirements
    1. Shop-A inventory link
    2. Shop-B Inventory link
    4. Shop-B should migrate Shop-A inventory

### APIs
    * Shop-A: Create a mock API that reads Shop-A inventory from JSON file. 
                1) Get ShopA inventory (ShopAInventoryHandler)
    * Shop-B: Create following APIs for shop be that do CRUD Opentions in SQLite
                1) Product Create and migrate (MigrateInventory)
                2) Product ListView  (ListProducts)

### Database Schema
    * Product_id
    * Product_title
    * Product_quantity
    * Product_price

### Code Directory Tree
* Note shop_b.db will be created when code will run
```
├── Readme.md
├── api
│   ├── shopA_api.go
│   └── shopB_api.go
├── data
│   └── shop_A_inventory.json
├── go.mod
├── go.sum
├── main.go
├── models
│   └── product.go
├── repository
│   ├── product_repository.go
│   └── product_repository_test.go
├── services
│   └── migration_service.go
└── shop_b.db
```
### How to Run
Clone the repo and navigate to workspace where main.go is present and run
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go run main.go
```

### How to Test
    1) Open browser
    2) paste http://127.0.0.1:8090/shopa/inventory on one tab should show ShopA inventory
    3) paste http://127.0.0.1:8090/shopb/inventory On 2nd tab should show empty inventory
    4) paste http://127.0.0.1:8090/shopb/migrate?shopAURL=http://localhost:8090/shopa/inventory on 3rd tab to migrate the from shop A - Shop B
    5) hop on to 2nd tab and refresh Shop A inventory is migrated to Shop B

