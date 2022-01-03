package roulette

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RequestPayload struct {
	UserID        string `json:"user_id"`
	CorrelationId string `json:"correlation_id"`
	Bets          []Bet  `json:"bets"`
}

type Bet struct {
	ID   string `json:"id"`
	Size int    `json:"size"`
	Type string `json:"type"`
}

type ResponsePayload struct {
	CorrelationId string `json:"correlation_id"`
	WinningNumber int    `json:"winning_number"`
	Winnings      int    `json:"winnings"`
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
