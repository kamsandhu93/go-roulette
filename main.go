package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/kamsandhu93/go-roulette/roulette"
	"os"
)

func setupRouter() *gin.Engine {
	router := gin.New() // without logger and recovery middleware

	router.GET("/health", func(c *gin.Context) {
		c.String(200, "ok")
	})
	router.POST("/v1/roulette", roulette.PostHandler)

	return router
}

func main() {
	router := setupRouter()

	port := getEnv("PORT", "8080")
	host := getEnv("HOST", "localhost")

	err := router.Run(host + ":" + port)
	if err != nil {
		panic(err)
	}
}

func getEnv(key, _default string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return _default
}
