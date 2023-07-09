package handlers

import (
	"bestwallet/services"

	"gorm.io/gorm"
)

type Handlers struct {
	db                 *gorm.DB
	accountService     *services.AccountService
	transactionService *services.TransactionService
}

// NewHandlers - Initializes and returns the Handlers
func NewHandlers(db *gorm.DB) Handlers {
	return Handlers{
		db:                 db,
		accountService:     services.NewAccountService(db),
		transactionService: services.NewTransactionService(db),
	}
}
