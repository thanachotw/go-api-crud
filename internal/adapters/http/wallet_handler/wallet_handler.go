package wallet_handler

import (
	"net/http"
	"strconv"

	"go-wallet-api/internal/core/domain"
	"go-wallet-api/internal/core/port"
	appresponse "go-wallet-api/pkg/appresponse"

	"github.com/gin-gonic/gin"
)

// SwaggerResponse represents a generic response for Swagger documentation.
type SwaggerResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type WalletHandler struct {
	Service port.WalletService
}

func NewWalletHandler(s port.WalletService) *WalletHandler {
	return &WalletHandler{Service: s}
}

// CreateWallet godoc
// @Summary Create a new wallet
// @Description Create a new wallet with the input payload
// @Tags wallets
// @Accept  json
// @Produce  json
// @Param wallet body CreateWalletRequest true "Create Wallet"
// @Success 200 {object} SwaggerResponse{data=WalletResponse}
// @Failure 400 {object} SwaggerResponse
// @Failure 500 {object} SwaggerResponse
// @Router /api/v1/wallets [post]
func (h *WalletHandler) CreateWallet(c *gin.Context) {
	var req CreateWalletRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		appresponse.ReponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	wallet, err := h.Service.CreateWallet(&domain.CreateWalletInput{
		OwnerID:  req.OwnerID,
		Currency: req.Currency,
	})
	if err != nil {
		appresponse.HandlerErrorResponse(c, err)
		return
	}
	appresponse.ResponseSuccess(c, appresponse.SuccessMessage, ToWalletResponse(wallet))
}

// GetWallet godoc
// @Summary Get a wallet
// @Description Get a wallet by owner ID
// @Tags wallets
// @Produce  json
// @Param owner_id path int true "Owner ID"
// @Success 200 {object} SwaggerResponse{data=WalletResponse}
// @Failure 400 {object} SwaggerResponse
// @Failure 404 {object} SwaggerResponse
// @Failure 500 {object} SwaggerResponse
// @Router /api/v1/wallets/{owner_id} [get]
func (h *WalletHandler) GetWallet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("owner_id"))
	wallet, err := h.Service.GetWallet(uint(id))
	if err != nil {
		appresponse.HandlerErrorResponse(c, err)
		return
	}
	appresponse.ResponseSuccess(c, appresponse.SuccessMessage, ToWalletResponse(wallet))
}

// UpdateWallet godoc
// @Summary Update a wallet
// @Description Update a wallet with the input payload
// @Tags wallets
// @Accept  json
// @Produce  json
// @Param owner_id path int true "Owner ID"
// @Param wallet body UpdateWalletRequest true "Update Wallet"
// @Success 200 {object} SwaggerResponse{data=WalletResponse}
// @Failure 400 {object} SwaggerResponse
// @Failure 404 {object} SwaggerResponse
// @Failure 500 {object} SwaggerResponse
// @Router /api/v1/wallets/{owner_id} [put]
func (h *WalletHandler) UpdateWallet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("owner_id"))
	var req UpdateWalletRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		appresponse.ReponseError(c, http.StatusBadRequest, err.Error())
		return
	}
	wallet, err := h.Service.UpdateWallet(&domain.UpdateWalletInput{
		OwnerID:  uint(id),
		Balance:  req.Balance,
		Currency: req.Currency,
		Status:   req.Status,
	})

	if err != nil {
		appresponse.HandlerErrorResponse(c, err)
		return
	}
	appresponse.ResponseSuccess(c, appresponse.SuccessMessage, ToWalletResponse(wallet))
}

// DeleteWallet godoc
// @Summary Delete a wallet
// @Description Delete a wallet by owner ID
// @Tags wallets
// @Produce  json
// @Param owner_id path int true "Owner ID"
// @Success 200 {object} SwaggerResponse
// @Failure 400 {object} SwaggerResponse
// @Failure 404 {object} SwaggerResponse
// @Failure 500 {object} SwaggerResponse
// @Router /api/v1/wallets/{owner_id} [delete]
func (h *WalletHandler) DeleteWallet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("owner_id"))
	err := h.Service.DeleteWallet(uint(id))
	if err != nil {
		appresponse.HandlerErrorResponse(c, err)
		return
	}
	appresponse.ResponseSuccess(c, appresponse.SuccessMessage, nil)
}
