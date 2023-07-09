package setup

import (
	"bestwallet/api/handlers"
	"bestwallet/api/middleware"
	"bestwallet/config"
	"bestwallet/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SetupDatabaseConnection - Sets up a database connection using gorm and performs automatic migrations.
func SetupDatabaseConnection() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.GetConfig().GetDBConnectionString()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Run the migrations for account and transaction
	db.AutoMigrate(&models.Account{}, &models.Transaction{})
	return db, nil
}

// SetupRouter - Configures and returns a Gin router with various API endpoints, along with authentication and verification middleware.
func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	handlers := handlers.NewHandlers(db)
	v1WithAuthentication := router.Group("/api/v1")
	v1 := router.Group("/api/v1")
	v1WithAuthentication.Use(middleware.HMACAuth())
	{
		v1WithAuthentication.POST("/account", middleware.KYCVerification(), handlers.CreateAccount())         // Create new account
		v1.GET("/account/:id", handlers.GetAccount())                                                         // Get account information
		v1WithAuthentication.PUT("/account/:id/deposit", handlers.DepositFunds())                             // Deposit funds
		v1WithAuthentication.PUT("/account/:id/withdraw", handlers.WithdrawFunds())                           // Withdraw funds
		v1WithAuthentication.POST("/transaction", middleware.KYTVerification(), handlers.CreateTransaction()) // Create a new transaction
	}

	return router
}
