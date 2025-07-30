package wallet

import (
	"go-wallet-api/internal/core/domain"

	"gorm.io/gorm"
)

type WalletRepo struct {
	DB *gorm.DB
}

func NewWalletRepo(db *gorm.DB) *WalletRepo {
	return &WalletRepo{DB: db}
}

func (r *WalletRepo) Create(walletModel *Wallet) error {
	return r.DB.Table(WalletTableName).Create(walletModel).Error
}

func (r *WalletRepo) GetByID(ownerId uint) (*domain.Wallet, error) {
	var walletModel Wallet
	result := r.DB.Table(WalletTableName).First(&walletModel, "owner_id = ?", ownerId)
	if result.Error != nil {
		return nil, result.Error
	}
	return walletModel.toDomain(), nil
}

func (r *WalletRepo) Update(walletModel *Wallet) error {
	return r.DB.Table(WalletTableName).Model(&Wallet{}).Where("owner_id = ?", walletModel.OwnerID).Updates(walletModel).Error
}

func (r *WalletRepo) Delete(ownerId uint) error {
	return r.DB.Table(WalletTableName).Where("owner_id = ?", ownerId).Delete(&Wallet{}).Error
}
