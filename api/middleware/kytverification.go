package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

// KYTVerification - simulates a KYT check with a delay of 60 seconds
func KYTVerification() gin.HandlerFunc {
	return func(c *gin.Context) {
		time.Sleep(60 * time.Second)
		c.Next()
	}
}
