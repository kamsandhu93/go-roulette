module gitlab.com/kamsandhu93/go-roulette

go 1.17

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/stretchr/testify v1.4.0
)

require github.com/go-playground/validator/v10 v10.4.1

replace gitlab.com/kamsandhu93/go-roulette/middleware v0.0.0 => ./middleware

replace gitlab.com/kamsandhu93/go-roulette/roulette v0.0.0 => ./roulette
