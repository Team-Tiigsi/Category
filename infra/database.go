package infra

import (
	"log"

	"github.com/Hamse/final_project/Back_end/Back_end/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectedDB() {
	dsn := "host=localhost user=postgres port=5432 password=123456 dbname=InventoryManagement sslmode=disable"
	db, err := gorm .Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect Database: ❌", err)
	}

	// AutoMigrate with error check
	err = db.AutoMigrate(
		&models.Category{},
	&models.Product{},
	)
	if err != nil {
		log.Fatal("failed to Auto migrate into database: ", err)
	}

	DB = db
	log.Println("✔️ Database connected and migrated successfully")
}