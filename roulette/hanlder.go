package roulette

import (
	"github.com/gin-gonic/gin"
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

func PostHandler(c *gin.Context) {
	//TODO
	return
}
