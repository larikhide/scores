package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SpyPlayerStore struct {
	store    map[string]int
	winCalls []string
}

func (s *SpyPlayerStore) GetPlayerScore(name string) int {
	score := s.store[name]
	return score
}

func (s *SpyPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	store := SpyPlayerStore{
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
	store := &SpyPlayerStore{
		store: map[string]int{},
	}
	server := &PlayerServer{
		store: store,
	}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		req := newPostWinRequest("Floyd")
		resp := httptest.NewRecorder()

		// act
		server.ServeHTTP(resp, req)

		// assert
		assertStatusCode(t, resp.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}
	})
}

func newPostWinRequest(name string) *http.Request {
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
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
