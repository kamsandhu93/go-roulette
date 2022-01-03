package main

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gitlab.com/kamsandhu93/go-roulette/roulette"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "ok", w.Body.String())
}

// Critical Path tests
func TestRoulleteStraight(t *testing.T) {
	router := setupRouter()

	bet := roulette.Bet{
		ID:   "1",
		Size: 1,
		Type: "17",
	}
	mockWinningNumber := 9
	expectedWinnings := 100

	reqBody, err := json.Marshal(roulette.RequestPayload{
		UserID:        "1",
		CorrelationId: "1",
		Bets:          []roulette.Bet{bet},
	})

	if err != nil {
		t.Fatal("[ERROR] Unable to generate the input json for the test", err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/roulette", bytes.NewBuffer(reqBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var respBody roulette.ResponsePayload
	if err := json.Unmarshal(w.Body.Bytes(), &respBody); err != nil {
		t.Fatal("[ERROR] Unable to deserialize JSON response", err)
	}

	assert.Equal(t, expectedWinnings, respBody.Winnings)
	assert.Equal(t, mockWinningNumber, respBody.WinningNumber)

}
