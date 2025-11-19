package middleware

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/genuicomm/api/pkg/config"
)

// CORS adalah middleware untuk menangani Cross-Origin Resource Sharing
func CORS() gin.HandlerFunc {
	corsConfig := config.GetCORSConfig()

	return func(c *gin.Context) {
		// Set CORS headers
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			for _, allowedOrigin := range corsConfig.AllowedOrigins {
				if allowedOrigin == "*" || allowedOrigin == origin {
					c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
					break
				}
			}
		}

		if corsConfig.AllowCredentials {
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		if len(corsConfig.AllowedHeaders) > 0 {
			c.Writer.Header().Set("Access-Control-Allow-Headers", strings.Join(corsConfig.AllowedHeaders, ","))
		}

		if len(corsConfig.AllowedMethods) > 0 {
			c.Writer.Header().Set("Access-Control-Allow-Methods", strings.Join(corsConfig.AllowedMethods, ","))
		}

		if corsConfig.MaxAge > 0 {
			c.Writer.Header().Set("Access-Control-Max-Age", strconv.Itoa(corsConfig.MaxAge))
		}

		// Handle preflight request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
