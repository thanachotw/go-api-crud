package domain

import (
	"time"
)

type Wallet struct {
	ID        uint
	OwnerID   uint
	Balance   float64
	Currency  string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateWalletInput struct {
	OwnerID  uint
	Currency string
}

type UpdateWalletInput struct {
	OwnerID  uint
	Balance  float64
	Currency string
	Status   string
}
