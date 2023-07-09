package handlers

import (
	"bestwallet/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTransaction - creates a new transaction
func (h *Handlers) CreateTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var transaction models.Transaction

		if err := c.ShouldBindJSON(&transaction); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := h.transactionService.CreateTransaction(&transaction)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, transaction)
	}
}
