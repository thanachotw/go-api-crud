package port

import (
	"go-wallet-api/internal/adapters/repository/wallet"
	"go-wallet-api/internal/core/domain"
)

type WalletRepository interface {
	Create(wallet *wallet.Wallet) error
	GetByID(ownerId uint) (*domain.Wallet, error)
	Update(wallet *wallet.Wallet) error
	Delete(id uint) error
}
