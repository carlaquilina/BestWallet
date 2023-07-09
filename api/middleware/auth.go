package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

const secret = "your_secret_key"

// HMACAuth - Creates a Gin middleware that performs HMAC authentication for incoming requests.
func HMACAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Read the Body content
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		// Restore the io.ReadCloser to its original state
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		// Use the method, URL, and body as the message
		message := c.Request.Method + c.Request.URL.String() + string(bodyBytes)

		// Continue with HMAC
		auth := c.GetHeader("Authorization")
		mac := hmac.New(sha256.New, []byte(secret))
		mac.Write([]byte(message))
		expectedMAC := hex.EncodeToString(mac.Sum(nil))

		if auth != expectedMAC {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		c.Next()
	}
}
