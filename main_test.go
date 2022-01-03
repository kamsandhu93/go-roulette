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
		Type: "straight",
	}

	reqBody, err := json.Marshal(roulette.RequestPayload{
		UserID:        "1",
		CorrelationId: "1",
		Bets:          []roulette.Bet{bet},
	})

	if err != nil {
		t.Fatalf("[ERROR] Unable to generate the input json for the test")
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/todo", bytes.NewBuffer(reqBody))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

}
