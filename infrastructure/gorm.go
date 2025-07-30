package infrastructure

import (
	"fmt"
	"go-wallet-api/config"
	"go-wallet-api/internal/adapters/repository/wallet"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitGormDB(config config.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		config.Host,
		config.User,
		config.Password,
		config.Name,
		config.Port,
	)
	log.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection established.")
	return db
}

func Migrate(db *gorm.DB) {
	log.Println("Running database migrations...")
	err := db.AutoMigrate(&wallet.Wallet{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migration completed.")
}
