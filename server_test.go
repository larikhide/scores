package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		// arrange
		req := makeGetScoreRequest("Pepper")
		resp := httptest.NewRecorder()

		// act
		PlayerServer(resp, req)

		//assert
		assertResponseBody(t, resp.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		// arrange
		req := makeGetScoreRequest("Floyd")
		resp := httptest.NewRecorder()

		// act
		PlayerServer(resp, req)

		// assert
		assertResponseBody(t, resp.Body.String(), "10")
	})
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got: %v\nwant: %v", got, want)
	}
}

func makeGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, name, nil)
	return req
}
