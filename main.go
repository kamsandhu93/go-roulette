package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gitlab.com/kamsandhu93/go-roulette/middleware"
	"gitlab.com/kamsandhu93/go-roulette/roulette"
	"os"
)

// Dependency injection of the spin wheel function used by the roulette route,
// allows tests to mock the winning number
func setupRouter(spinWheelFunc roulette.SpinWheelFunc) *gin.Engine {
	router := gin.New() // without logger and recovery middleware

	router.Use(gin.Logger())
	router.Use(middleware.Logger()) //extra logging
	router.Use(middleware.Auth())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.GET("/health", func(c *gin.Context) {
		c.String(200, "ok")
	})
	router.POST("/v1/roulette", func(context *gin.Context) {
		roulette.PostHandler(context, spinWheelFunc)

	})
	return router
}

func SetUpBindValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("validBetType", roulette.ValidBetType)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	SetUpBindValidation()
	router := setupRouter(roulette.SpinWheel)

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
