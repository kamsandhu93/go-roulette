package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.New() // without logger and recovery middleware
	router.GET("/health", func(c *gin.Context) {
		c.String(200, "ok")
	})
	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
