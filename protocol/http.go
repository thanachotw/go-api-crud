package protocol

import (
	"go-wallet-api/config"
	_ "go-wallet-api/docs"
	"go-wallet-api/infrastructure"
	walletH "go-wallet-api/internal/adapters/http/wallet_handler"
	"go-wallet-api/internal/adapters/repository/wallet"
	service "go-wallet-api/internal/core/service/wallet"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Wallet API
// @version 1.0
// @description This is a sample server for a wallet API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @schemes https
// @schemes https
func ServeREST() error {
	cfg := config.InitConfig()
	db := infrastructure.InitGormDB(cfg.UserWalletDB)
	infrastructure.Migrate(db)

	// Seed the database with initial data
	infrastructure.SeedDatabase(db)

	walletRepo := wallet.NewWalletRepo(db)
	walletSvc := service.NewWalletService(walletRepo)
	walletHandler := walletH.NewWalletHandler(walletSvc)

	r := gin.Default()
	api := r.Group("/api")
	apiV1 := api.Group("/v1")
	// Register wallet routes
	apiV1.POST("/wallets", walletHandler.CreateWallet)
	apiV1.GET("/wallets/:owner_id", walletHandler.GetWallet)
	apiV1.PUT("/wallets/:owner_id", walletHandler.UpdateWallet)
	apiV1.DELETE("/wallets/:owner_id", walletHandler.DeleteWallet)

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check endpoint
	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return r.Run(":" + cfg.AppsConfig.Port)
}
