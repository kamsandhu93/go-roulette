package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)
		status := c.Writer.Status()

		log.Printf("Latency=%d Status=%d", latency, status)
	}
}
