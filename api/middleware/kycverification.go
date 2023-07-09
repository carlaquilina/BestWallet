package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

// KYCVerification - simulates a KYC check with a delay of 60 seconds
func KYCVerification() gin.HandlerFunc {
	return func(c *gin.Context) {
		time.Sleep(60 * time.Second)
		c.Next()
	}
}
