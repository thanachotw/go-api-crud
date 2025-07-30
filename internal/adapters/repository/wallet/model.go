package wallet

import (
	"go-wallet-api/internal/core/domain"
	"time"
)

var (
	WalletTableName = "wallets"
)

// Wallet is the GORM model for the wallets table.
// It is kept separate from the core domain model.
type Wallet struct {
	ID        uint `gorm:"primaryKey"`
	OwnerID   uint
	Balance   float64
	Currency  string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// toDomain converts the repository's Wallet model to the core domain.Wallet model.
func (w *Wallet) toDomain() *domain.Wallet {
	return &domain.Wallet{
		ID:        w.ID,
		OwnerID:   w.OwnerID,
		Balance:   w.Balance,
		Currency:  w.Currency,
		Status:    w.Status,
		CreatedAt: w.CreatedAt,
		UpdatedAt: w.UpdatedAt,
	}
}
