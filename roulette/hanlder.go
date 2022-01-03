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

	//TODO stubs for now
	winningNumber := spinWheelFunc()
	winnings := 100

	// Indented JSON used for debugging only - production app should call c.JSON
	c.IndentedJSON(http.StatusOK, ResponsePayload{
		CorrelationId: requestPayload.CorrelationId,
		WinningNumber: winningNumber,
		Winnings:      winnings,
	})

	return
}

type SpinWheelFunc func() int

func SpinWheel() int {
	return 9
	//TODO hardcoded for now, will eventually return a random number between 0 and 36
}
