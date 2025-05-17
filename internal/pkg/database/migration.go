package database

import (
	"log"

	"github.com/colinjuang/shop-go/internal/model"
)

// AutoMigrate performs database migrations
func AutoMigrate() error {
	log.Println("Starting database migration...")

	// Add all models here for auto migration
	err := DB.AutoMigrate(
		&model.Banner{},
		&model.CartItem{},
		&model.Order{},
		&model.OrderItem{},
		&model.Promotion{},
		// Add other models here
	)
	if err != nil {
		log.Printf("Failed to perform database migration: %v", err)
		return err
	}

	log.Println("Database migration completed successfully")
	return nil
}
