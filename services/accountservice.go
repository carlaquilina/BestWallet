package services

import (
	"bestwallet/models"
	"errors"

	"gorm.io/gorm"
)

// AccountService handles operations related to accounts
type AccountService struct {
	db *gorm.DB
}

// NewAccountService creates a new AccountService
func NewAccountService(db *gorm.DB) *AccountService {
	return &AccountService{
		db: db,
	}
}

// CreateAccount creates a new account
func (service *AccountService) CreateAccount(account *models.Account) error {
	return service.db.Create(account).Error
}

// GetAccount retrieves an account by ID
func (service *AccountService) GetAccount(id int) (*models.Account, error) {
	var account models.Account
	err := service.db.First(&account, id).Error
	return &account, err
}

// Deposit adds funds into an account
func (service *AccountService) Deposit(id int, amount float64) error {
	var account models.Account
	err := service.db.First(&account, id).Error
	if err != nil {
		return err
	}

	account.Balance += amount
	return service.db.Save(&account).Error
}

// Withdraw subtracts funds from an account
func (service *AccountService) Withdraw(id int, amount float64) error {
	var account models.Account
	err := service.db.First(&account, id).Error
	if err != nil {
		return err
	}

	if account.Balance < amount {
		return errors.New("insufficient funds")
	}

	account.Balance -= amount
	return service.db.Save(&account).Error
}
