package wallet

import (
	"errors"
	"go-wallet-api/internal/core/domain"
	"go-wallet-api/internal/core/port"
	appresponse "go-wallet-api/pkg/appresponse"
	"strings"
	"time"

	"gorm.io/gorm"
)

type WalletService struct {
	Repo port.WalletRepository
}

func NewWalletService(repo port.WalletRepository) port.WalletService {
	return &WalletService{Repo: repo}
}

func (s *WalletService) CreateWallet(input *domain.CreateWalletInput) (*domain.Wallet, error) {
	if input.OwnerID == 0 {
		return nil, appresponse.ErrOwnerIDInvalid
	}
	if input.Currency != "THB" && input.Currency != "USD" {
		return nil, appresponse.ErrUnsupportedCurrency
	}
	w := &domain.Wallet{
		OwnerID:   input.OwnerID,
		Balance:   0,
		Currency:  input.Currency,
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.Repo.Create(toWalletModel(w))
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (s *WalletService) GetWallet(id uint) (*domain.Wallet, error) {
	wallet, err := s.Repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appresponse.ErrNotFound
		}
		return nil, appresponse.ErrInternalServer
	}
	if wallet == nil {
		return nil, appresponse.ErrNotFound
	}
	return wallet, nil
}

func (s *WalletService) UpdateWallet(input *domain.UpdateWalletInput) (*domain.Wallet, error) {
	w, err := s.Repo.GetByID(input.OwnerID)
	if err != nil {
		return nil, err
	}
	if input.Currency != "" {
		if input.Currency != "THB" && input.Currency != "USD" {
			return nil, errors.New("unsupported currency")
		}
		w.Currency = input.Currency
	}
	if input.Status != "" {
		status := strings.ToLower(input.Status)
		if status != "active" && status != "disabled" {
			return nil, errors.New("unsupported status")
		}
		w.Status = status
	}
	w.UpdatedAt = time.Now()
	w.Balance = input.Balance
	err = s.Repo.Update(toWalletModel(w))
	return w, err
}

func (s *WalletService) DeleteWallet(id uint) error {
	return s.Repo.Delete(id)
}
