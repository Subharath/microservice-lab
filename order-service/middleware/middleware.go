package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

// CORSConfig holds CORS configuration
type CORSConfig struct {
	AllowOrigin string
}

// CORS returns a CORS middleware
func CORS(config CORSConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", config.AllowOrigin)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "86400")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Logger returns a logging middleware
func Logger(serviceName string) gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return formatLog(serviceName, param)
	})
}

func formatLog(serviceName string, param gin.LogFormatterParams) string {
	// JSON structured logging
	return `{"timestamp":"` + param.TimeStamp.Format(time.RFC3339) +
		`","service":"` + serviceName +
		`","level":"info` +
		`","method":"` + param.Method +
		`","path":"` + param.Path +
		`","status":` + string(rune(param.StatusCode+'0')) +
		`","latency":"` + param.Latency.String() +
		`","ip":"` + param.ClientIP +
		`"}` + "\n"
}

// Recovery returns a recovery middleware that recovers from panics
func Recovery() gin.HandlerFunc {
	return gin.Recovery()
}
