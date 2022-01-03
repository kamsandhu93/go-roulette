package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gitlab.com/kamsandhu93/go-roulette/roulette"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthRoute(t *testing.T) {
	router := setupRouter(roulette.SpinWheel)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "ok", w.Body.String())
}

func TestRoulettePostCriticalPathTable(t *testing.T) {
	tests := map[string]struct {
		inputBets         []roulette.Bet
		mockWinningNumber int
		expectedWinnings  int
	}{
		"single_straight_win": {mockWinningNumber: 0, expectedWinnings: 108, inputBets: []roulette.Bet{
			{ID: "0", Size: 3, Type: "0"}}},
		"single_straight_loss": {mockWinningNumber: 1, expectedWinnings: 0, inputBets: []roulette.Bet{
			{ID: "0", Size: 3, Type: "0"}}},
		"multiple_straight_win": {mockWinningNumber: 0, expectedWinnings: 108, inputBets: []roulette.Bet{
			{ID: "0", Size: 3, Type: "0"},
			{ID: "0", Size: 10, Type: "5"},
			{ID: "0", Size: 4, Type: "30"}}},
		"multiple_straight_loss": {mockWinningNumber: 1, expectedWinnings: 0, inputBets: []roulette.Bet{
			{ID: "0", Size: 3, Type: "0"},
			{ID: "0", Size: 10, Type: "5"},
			{ID: "0", Size: 4, Type: "30"}}},
		"single_red_win": {mockWinningNumber: 1, expectedWinnings: 6, inputBets: []roulette.Bet{
			{ID: "0", Size: 3, Type: "red"}}},
		"single_red_loss": {mockWinningNumber: 2, expectedWinnings: 0, inputBets: []roulette.Bet{
			{ID: "0", Size: 3, Type: "red"}}},
		"single_odd_win": {mockWinningNumber: 19, expectedWinnings: 6, inputBets: []roulette.Bet{
			{ID: "0", Size: 3, Type: "odd"}}},
		"single_odd_loss": {mockWinningNumber: 16, expectedWinnings: 0, inputBets: []roulette.Bet{
			{ID: "0", Size: 3, Type: "odd"}}},
		"single_high_win": {mockWinningNumber: 36, expectedWinnings: 6, inputBets: []roulette.Bet{
			{ID: "0", Size: 3, Type: "high"}}},
		"single_low_loss": {mockWinningNumber: 1, expectedWinnings: 0, inputBets: []roulette.Bet{
			{ID: "0", Size: 3, Type: "high"}}},
		"single_square_win": {mockWinningNumber: 1, expectedWinnings: 27, inputBets: []roulette.Bet{
			{ID: "0", Size: 3, Type: "1-2-4-5"}}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			fmt.Println("--- Test case:", name, "---")

			reqBody, err := json.Marshal(roulette.RequestPayload{
				UserID:        "1",
				CorrelationId: "1",
				Bets:          tc.inputBets,
			})

			if err != nil {
				t.Fatal("[ERROR] Unable to generate the input json for the test", err)
			}

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/v1/roulette", bytes.NewBuffer(reqBody))

			router := setupRouter(func() int {
				return tc.mockWinningNumber
			})
			router.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code, "Expected http status code not equal actual code")

			var respBody roulette.ResponsePayload
			if err := json.Unmarshal(w.Body.Bytes(), &respBody); err != nil {
				t.Fatal("[ERROR] Unable to deserialize JSON response", err)
			}

			assert.Equal(t, tc.expectedWinnings, respBody.Winnings, "Expected winnings not equal actual winnings")
			assert.Equal(t, tc.mockWinningNumber, respBody.WinningNumber, "Expected winning number not equal actual winning number")
		})
	}
}
