package roulette

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type RequestPayload struct {
	UserID        string `json:"user_id" binding:"required"`
	CorrelationId string `json:"correlation_id" binding:"required"`
	Bets          []Bet  `json:"bets" binding:"required,dive"`
}

type Bet struct {
	ID   string `json:"id" binding:"required"`
	Size int    `json:"size" binding:"required"`
	Type string `json:"type" binding:"required,validBetType"`
}

type ResponsePayload struct {
	CorrelationId string `json:"correlation_id"`
	WinningNumber int    `json:"winning_number"`
	Winnings      int    `json:"winnings"`
}

var ValidBetType validator.Func = func(fl validator.FieldLevel) bool {
	betType, ok := fl.Field().Interface().(string)
	if ok {
		if _, _ok := betTypesMap[betType]; _ok {
			return true
		}
	}
	return false
}

func PostHandler(c *gin.Context, spinWheelFunc SpinWheelFunc) {
	var requestPayload RequestPayload
	if err := c.BindJSON(&requestPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	winningNumber, winnings := postController(requestPayload.Bets, spinWheelFunc)

	// Indented JSON used for debugging only - a production app would call c.JSON
	c.IndentedJSON(http.StatusOK, ResponsePayload{
		CorrelationId: requestPayload.CorrelationId,
		WinningNumber: winningNumber,
		Winnings:      winnings,
	})

	return
}
