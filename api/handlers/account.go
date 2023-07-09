package handlers

import (
	"bestwallet/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateAccount - creates a new account
func (h *Handlers) CreateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var account models.Account

		if err := c.ShouldBindJSON(&account); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := h.accountService.CreateAccount(&account)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, account)
	}
}

// GetAccount - retrieves an account by ID
func (h *Handlers) GetAccount() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, _ := strconv.Atoi(c.Param("id"))

		account, err := h.accountService.GetAccount(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, account)
	}
}

// DepositFunds - deposits funds into an account
func (h *Handlers) DepositFunds() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, _ := strconv.Atoi(c.Param("id"))
		var payload struct {
			Amount float64 `json:"amount"`
		}

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := h.accountService.Deposit(id, payload.Amount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}

// WithdrawFunds - withdraws funds from an account
func (h *Handlers) WithdrawFunds() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, _ := strconv.Atoi(c.Param("id"))
		var payload struct {
			Amount float64 `json:"amount"`
		}

		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := h.accountService.Withdraw(id, payload.Amount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}
