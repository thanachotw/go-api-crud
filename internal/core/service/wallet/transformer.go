package wallet

import (
	repo "go-wallet-api/internal/adapters/repository/wallet"
	domain "go-wallet-api/internal/core/domain"
)

// fromDomain converts a core domain.Wallet model to the repository's Wallet model.
func toWalletModel(d *domain.Wallet) *repo.Wallet {
	return &repo.Wallet{
		ID:        d.ID,
		OwnerID:   d.OwnerID,
		Balance:   d.Balance,
		Currency:  d.Currency,
		Status:    d.Status,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}
