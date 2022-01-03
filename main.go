package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.New() // without logger and recovery middleware

	router.GET("/health", func(c *gin.Context) {
		c.String(200, "ok")
	})

	return router
}

func main() {
	router := setupRouter()
	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}
