package services

import (
	"bestwallet/models"
	"fmt"

	"gorm.io/gorm"
)

// TransactionService - handles operations related to transactions
type TransactionService struct {
	db *gorm.DB
}

// NewTransactionService - creates a new TransactionService
func NewTransactionService(db *gorm.DB) *TransactionService {
	return &TransactionService{
		db: db,
	}
}

// CreateTransaction - creates a new transaction
func (service *TransactionService) CreateTransaction(transaction *models.Transaction) error {
	fromAccount := &models.Account{}
	toAccount := &models.Account{}

	err := service.db.First(fromAccount, transaction.FromAccountID).Error
	if err != nil {
		return fmt.Errorf("could not find the account to transfer from: %v", err)
	}

	err = service.db.First(toAccount, transaction.ToAccountID).Error
	if err != nil {
		return fmt.Errorf("could not find the account to transfer to: %v", err)
	}

	if fromAccount.Balance < transaction.Amount {
		return fmt.Errorf("insufficient balance")
	}

	fromAccount.Balance -= transaction.Amount
	toAccount.Balance += transaction.Amount

	tx := service.db.Begin()

	err = tx.Save(fromAccount).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("could not deduct amount from account: %v", err)
	}

	err = tx.Save(toAccount).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("could not add amount to account: %v", err)
	}

	err = tx.Create(transaction).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("could not create transaction: %v", err)
	}

	tx.Commit()

	return nil
}
