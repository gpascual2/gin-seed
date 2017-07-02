// Package middleware --> CORS provider
package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//CORSMiddleware returns a gin.HandlerFunc (middleware) that setups CORS on HTTP Headers
func CORSMiddleware(config *viper.Viper) gin.HandlerFunc {
	OriginKey := "Origin"
	var origins []string
	origins = config.GetStringSlice("cors.allow_origin")

	return func(c *gin.Context) {
		// Read the Origin header from the HTTP request
		currentOrigin := c.Request.Header.Get(OriginKey)
		c.Writer.Header().Add("Vary", OriginKey)
		if currentOrigin == "" {
			return
		}
		if matchOrigin(currentOrigin, origins) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", currentOrigin)
			c.Writer.Header().Set("Access-Control-Max-Age", "86400")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
			c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

			if c.Request.Method == "OPTIONS" {
				fmt.Println("OPTIONS")
				c.AbortWithStatus(200)
			} else {
				c.Next()
			}
		}
	}
}

// Case-sensitive match of origin header
func matchOrigin(currentOrigin string, allowedOrigins []string) bool {
	for _, value := range allowedOrigins {
		if value == currentOrigin {
			return true
		}
	}
	return false
}
