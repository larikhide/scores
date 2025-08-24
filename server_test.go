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
	score := s.store[name]
	return score
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
		assertStatusCode(t, resp.Code, http.StatusOK)
		assertResponseBody(t, resp.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		// arrange
		req := makeGetScoreRequest("Floyd")
		resp := httptest.NewRecorder()

		// act
		server.ServeHTTP(resp, req)

		// assert
		assertStatusCode(t, resp.Code, http.StatusOK)
		assertResponseBody(t, resp.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		// arrange
		req := makeGetScoreRequest("Apollo")
		resp := httptest.NewRecorder()

		//act
		server.ServeHTTP(resp, req)

		// assert
		assertStatusCode(t, resp.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	// arrange
	store := &StubPlayerStore{
		store: map[string]int{},
	}
	server := &PlayerServer{
		store: store,
	}

	t.Run("It returns accepted on POST", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/players/Pepper", nil)
		resp := httptest.NewRecorder()

		// act
		server.ServeHTTP(resp, req)

		// assert
		assertStatusCode(t, resp.Code, http.StatusAccepted)
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

func assertStatusCode(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got: %d\nwant: %d", got, want)
	}
}
