package infrastructure

import (
	repo "go-wallet-api/internal/adapters/repository/wallet"
	"log"

	"gorm.io/gorm"
)

// SeedDatabase populates the database with initial data.
// It uses FirstOrCreate to ensure that data is not duplicated on subsequent runs.
func SeedDatabase(db *gorm.DB) {
	wallets := []repo.Wallet{
		{OwnerID: 101, Balance: 1500.00, Currency: "USD", Status: "active"},
		{OwnerID: 102, Balance: 250.75, Currency: "EUR", Status: "active"},
		{OwnerID: 103, Balance: 9999.99, Currency: "JPY", Status: "inactive"},
	}

	for _, wallet := range wallets {
		// Check if a wallet with the same OwnerID already exists.
		// If it does, do nothing. If it doesn't, create it.
		result := db.FirstOrCreate(&wallet, wallet)
		if result.Error != nil {
			log.Printf("Failed to seed wallet for owner %d: %v\n", wallet.OwnerID, result.Error)
		}

		if result.RowsAffected > 0 {
			log.Printf("Seeded wallet for owner %d\n", wallet.OwnerID)
		}
	}

	log.Println("Database seeding completed.")
}
