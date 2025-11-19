package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger adalah middleware untuk logging request
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		// Log request
		fmt.Printf("[%s] %s %s %d %v\n",
			end.Format("2006/01/02 - 15:04:05"),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			latency,
		)
	}
}
