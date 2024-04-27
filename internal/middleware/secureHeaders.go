package middleware

import (
	"github.com/gin-gonic/gin"
)

// SecureHeadersMiddleware adds security headers to each response
func SecureHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.Header("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-XSS-Protection", "1; mode=block")

		c.Next()
	}
}
