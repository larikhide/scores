package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	store map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.store[name]
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		store: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		}}
	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		// arrange
		req := makeGetScoreRequest("Pepper")
		resp := httptest.NewRecorder()

		// act
		server.ServeHTTP(resp, req)

		//assert
		assertResponseBody(t, resp.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		// arrange
		req := makeGetScoreRequest("Floyd")
		resp := httptest.NewRecorder()

		// act
		server.ServeHTTP(resp, req)

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
