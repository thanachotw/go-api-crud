package wallet_handler

import (
	"go-wallet-api/internal/core/domain"
	"time"
)

type CreateWalletRequest struct {
	OwnerID  uint   `json:"owner_id"`
	Currency string `json:"currency"`
}

type UpdateWalletRequest struct {
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance"`
	Status   string  `json:"status"`
}

// WalletResponse represents a wallet for a user.
type WalletResponse struct {
	ID        uint      `json:"id"`
	OwnerID   uint      `json:"owner_id"`
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (r *WalletResponse) ToDomain() *domain.Wallet {
	return &domain.Wallet{
		ID:        r.ID,
		OwnerID:   r.OwnerID,
		Balance:   r.Balance,
		Currency:  r.Currency,
		Status:    r.Status,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

func ToWalletResponse(wallet *domain.Wallet) *WalletResponse {
	return &WalletResponse{
		ID:        wallet.ID,
		OwnerID:   wallet.OwnerID,
		Balance:   wallet.Balance,
		Currency:  wallet.Currency,
		Status:    wallet.Status,
		CreatedAt: wallet.CreatedAt,
		UpdatedAt: wallet.UpdatedAt,
	}
}
