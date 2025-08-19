package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		// arrange
		req, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		resp := httptest.NewRecorder()

		// act
		PlayerServer(resp, req)
		got := resp.Body.String()

		//assert
		want := "20"
		if got != want {
			t.Errorf("got: %v\nwant: %v", got, want)
		}
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		// arrange
		req, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		resp := httptest.NewRecorder()

		// act
		PlayerServer(resp, req)
		got := resp.Body.String()

		// assert
		want := "10"
		if got != want {
			t.Errorf("got: %v\nwant: %v", got, want)
		}
	})
}
