package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryStore()
	server := NewPlayerServer(store)
	player := "Pepper"

	t.Run("get player score", func(t *testing.T) {
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

		resp := httptest.NewRecorder()
		server.ServeHTTP(resp, makeGetScoreRequest(player))

		assertStatusCode(t, resp.Code, http.StatusOK)
		assertResponseBody(t, resp.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		resp := httptest.NewRecorder()
		req := makeLeagueRequest()
		want := []Player{
			{"Pepper", 3},
		}

		server.ServeHTTP(resp, req)
		got := getLeagueFromResponse(t, resp.Body)

		assertStatusCode(t, resp.Code, http.StatusOK)
		assertLeague(t, want, got)
	})
}
