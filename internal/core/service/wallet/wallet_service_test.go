package wallet_test

import (
	"errors"
	"testing"
	"time"

	walletRepo "go-wallet-api/internal/adapters/repository/wallet"
	"go-wallet-api/internal/core/domain"
	"go-wallet-api/internal/core/port"
	walletService "go-wallet-api/internal/core/service/wallet"
	appresponse "go-wallet-api/pkg/appresponse"
	"go-wallet-api/tests/mocks"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type WalletServiceTestSuite struct {
	suite.Suite
	mockWalletRepo *mocks.MockWalletRepository
	walletService  port.WalletService
}

func (s *WalletServiceTestSuite) SetupSuite() {
	s.mockWalletRepo = mocks.NewMockWalletRepository(s.T())
	s.walletService = walletService.NewWalletService(s.mockWalletRepo)
}

func (s *WalletServiceTestSuite) TearDownTest() {
	s.mockWalletRepo.Calls = nil
	s.mockWalletRepo.ExpectedCalls = nil
}

func (s *WalletServiceTestSuite) TestCreateWallet_Success() {
	input := &domain.CreateWalletInput{
		OwnerID:  1,
		Currency: "THB",
	}
	expectedWallet := &domain.Wallet{
		OwnerID:   1,
		Balance:   0,
		Currency:  "THB",
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	walletInput := &walletRepo.Wallet{
		OwnerID:   input.OwnerID,
		Currency:  input.Currency,
		Balance:   expectedWallet.Balance,
		Status:    expectedWallet.Status,
		CreatedAt: expectedWallet.CreatedAt,
		UpdatedAt: expectedWallet.UpdatedAt,
	}
	s.mockWalletRepo.On("Create", walletInput).Return(nil).Once()

	wallet, err := s.walletService.CreateWallet(input)

	s.Require().NoError(err)
	s.Require().Equal(expectedWallet.OwnerID, wallet.OwnerID)
	s.Require().Equal(expectedWallet.Currency, wallet.Currency)
	s.Require().Equal(expectedWallet.Balance, wallet.Balance)
	s.Require().Equal(expectedWallet.Status, wallet.Status)
	s.mockWalletRepo.AssertExpectations(s.T())
}

func (s *WalletServiceTestSuite) TestCreateWallet_InvalidOwnerID() {
	input := &domain.CreateWalletInput{
		OwnerID:  0,
		Currency: "THB",
	}

	wallet, err := s.walletService.CreateWallet(input)

	s.Require().Equal(appresponse.ErrOwnerIDInvalid, err)
	s.Require().Nil(wallet)
	s.mockWalletRepo.AssertNotCalled(s.T(), "Create")
}

func (s *WalletServiceTestSuite) TestCreateWallet_UnsupportedCurrency() {
	input := &domain.CreateWalletInput{
		OwnerID:  1,
		Currency: "XYZ",
	}

	wallet, err := s.walletService.CreateWallet(input)

	s.Require().Equal(appresponse.ErrUnsupportedCurrency, err)
	s.Require().Nil(wallet)
	s.mockWalletRepo.AssertNotCalled(s.T(), "Create")
}

func (s *WalletServiceTestSuite) TestCreateWallet_RepositoryError() {
	input := &domain.CreateWalletInput{
		OwnerID:  1,
		Currency: "THB",
	}
	repoError := errors.New("database error")
	walletInput := &walletRepo.Wallet{
		OwnerID:   input.OwnerID,
		Currency:  input.Currency,
		Status:    "active",
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	s.mockWalletRepo.On("Create", walletInput).Return(repoError).Once()

	wallet, err := s.walletService.CreateWallet(input)

	s.Require().Equal(repoError, err)
	s.Require().Nil(wallet)
	s.mockWalletRepo.AssertExpectations(s.T())
}

func (s *WalletServiceTestSuite) TestGetWallet_Success() {
	ownerID := uint(1)
	expectedWallet := &domain.Wallet{
		ID:       1,
		OwnerID:  ownerID,
		Balance:  100,
		Currency: "THB",
		Status:   "active",
	}

	s.mockWalletRepo.On("GetByID", ownerID).Return(expectedWallet, nil).Once()

	wallet, err := s.walletService.GetWallet(ownerID)

	s.Require().NoError(err)
	s.Require().Equal(expectedWallet, wallet)
	s.mockWalletRepo.AssertExpectations(s.T())
}

func (s *WalletServiceTestSuite) TestGetWallet_NotFound() {
	ownerID := uint(999)

	s.mockWalletRepo.On("GetByID", ownerID).Return(nil, gorm.ErrRecordNotFound).Once()

	wallet, err := s.walletService.GetWallet(ownerID)

	s.Require().Equal(appresponse.ErrNotFound, err)
	s.Require().Nil(wallet)
	s.mockWalletRepo.AssertExpectations(s.T())
}

func (s *WalletServiceTestSuite) TestGetWallet_RepositoryError() {
	ownerID := uint(1)
	repoError := errors.New("database connection error")

	s.mockWalletRepo.On("GetByID", ownerID).Return(nil, repoError).Once()

	wallet, err := s.walletService.GetWallet(ownerID)

	s.Require().Equal(appresponse.ErrInternalServer, err)
	s.Require().Nil(wallet)
	s.mockWalletRepo.AssertExpectations(s.T())
}

func (s *WalletServiceTestSuite) TestUpdateWallet_Success() {
	ownerID := uint(1)
	existingWallet := &domain.Wallet{
		ID:        1,
		OwnerID:   ownerID,
		Balance:   100,
		Currency:  "THB",
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	input := &domain.UpdateWalletInput{
		OwnerID:  ownerID,
		Currency: "USD",
		Status:   "disabled",
	}
	walletInput := &walletRepo.Wallet{
		ID:        existingWallet.ID,
		OwnerID:   input.OwnerID,
		Currency:  input.Currency,
		Status:    input.Status,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	s.mockWalletRepo.On("GetByID", ownerID).Return(existingWallet, nil).Once()
	s.mockWalletRepo.On("Update", walletInput).Return(nil).Once()

	updatedWallet, err := s.walletService.UpdateWallet(input)

	s.Require().NoError(err)
	s.Require().Equal(input.Currency, updatedWallet.Currency)
	s.Require().Equal(input.Status, updatedWallet.Status)
	s.Require().Equal(input.Balance, updatedWallet.Balance)
	s.mockWalletRepo.AssertExpectations(s.T())
}

func (s *WalletServiceTestSuite) TestUpdateWallet_GetByIDError() {
	ownerID := uint(1)
	repoError := errors.New("get by ID error")
	input := &domain.UpdateWalletInput{
		OwnerID: ownerID,
	}

	s.mockWalletRepo.On("GetByID", ownerID).Return(nil, repoError).Once()

	updatedWallet, err := s.walletService.UpdateWallet(input)

	s.Require().Equal(repoError, err)
	s.Require().Nil(updatedWallet)
	s.mockWalletRepo.AssertExpectations(s.T())
}

func (s *WalletServiceTestSuite) TestUpdateWallet_UnsupportedCurrency() {
	ownerID := uint(1)
	existingWallet := &domain.Wallet{
		ID:        1,
		OwnerID:   ownerID,
		Balance:   100,
		Currency:  "THB",
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	input := &domain.UpdateWalletInput{
		OwnerID:  ownerID,
		Currency: "XYZ",
	}

	s.mockWalletRepo.On("GetByID", ownerID).Return(existingWallet, nil).Once()

	updatedWallet, err := s.walletService.UpdateWallet(input)

	s.Require().Equal(errors.New("unsupported currency"), err)
	s.Require().Nil(updatedWallet)
	s.mockWalletRepo.AssertExpectations(s.T())
	s.mockWalletRepo.AssertNotCalled(s.T(), "Update")
}

func (s *WalletServiceTestSuite) TestUpdateWallet_UnsupportedStatus() {
	ownerID := uint(1)
	existingWallet := &domain.Wallet{
		ID:        1,
		OwnerID:   ownerID,
		Balance:   100,
		Currency:  "THB",
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	input := &domain.UpdateWalletInput{
		OwnerID: ownerID,
		Status:  "invalid",
	}

	s.mockWalletRepo.On("GetByID", ownerID).Return(existingWallet, nil).Once()

	updatedWallet, err := s.walletService.UpdateWallet(input)

	s.Require().Equal(errors.New("unsupported status"), err)
	s.Require().Nil(updatedWallet)
	s.mockWalletRepo.AssertExpectations(s.T())
	s.mockWalletRepo.AssertNotCalled(s.T(), "Update")
}

func (s *WalletServiceTestSuite) TestUpdateWallet_RepositoryError() {
	ownerID := uint(1)
	existingWallet := &domain.Wallet{
		ID:        1,
		OwnerID:   ownerID,
		Balance:   100,
		Currency:  "THB",
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	input := &domain.UpdateWalletInput{
		OwnerID:  ownerID,
		Currency: "USD",
	}
	repoError := errors.New("update failed")
	walletInput := &walletRepo.Wallet{
		ID:        existingWallet.ID,
		OwnerID:   input.OwnerID,
		Currency:  input.Currency,
		Status:    existingWallet.Status,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	s.mockWalletRepo.On("GetByID", ownerID).Return(existingWallet, nil).Once()
	s.mockWalletRepo.On("Update", walletInput).Return(repoError).Once()

	updatedWallet, err := s.walletService.UpdateWallet(input)

	s.Require().Equal(repoError, err)
	s.Require().Equal(existingWallet, updatedWallet) // Should return the original wallet on error
	s.mockWalletRepo.AssertExpectations(s.T())
}

func (s *WalletServiceTestSuite) TestDeleteWallet_Success() {
	ownerID := uint(1)

	s.mockWalletRepo.On("Delete", ownerID).Return(nil).Once()

	err := s.walletService.DeleteWallet(ownerID)

	s.Require().NoError(err)
	s.mockWalletRepo.AssertExpectations(s.T())
}

func (s *WalletServiceTestSuite) TestDeleteWallet_RepositoryError() {
	ownerID := uint(1)
	repoError := errors.New("delete failed")

	s.mockWalletRepo.On("Delete", ownerID).Return(repoError).Once()

	err := s.walletService.DeleteWallet(ownerID)

	s.Require().Equal(repoError, err)
	s.mockWalletRepo.AssertExpectations(s.T())
}

func TestWalletServiceTestSuite(t *testing.T) {
	suite.Run(t, new(WalletServiceTestSuite))
}
