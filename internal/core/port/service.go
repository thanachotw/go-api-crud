package port

import "go-wallet-api/internal/core/domain"

type WalletService interface {
	CreateWallet(input *domain.CreateWalletInput) (*domain.Wallet, error)
	GetWallet(id uint) (*domain.Wallet, error)
	UpdateWallet(input *domain.UpdateWalletInput) (*domain.Wallet, error)
	DeleteWallet(id uint) error
}
