package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Call to auth library...
		log.Println("[INFO] Auth success")
		c.Next()
	}
}
