package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type SpyPlayerStore struct {
	store    map[string]int
	winCalls []string
	league   []Player
}

// GetLeague implements playerStore.
func (s *SpyPlayerStore) GetLeague() []Player {
	return s.league
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
	server := NewPlayerServer(&store)

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
	server := NewPlayerServer(store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Floyd"
		req := newPostWinRequest(player)
		resp := httptest.NewRecorder()

		// act
		server.ServeHTTP(resp, req)

		// assert
		assertStatusCode(t, resp.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not score correct winner got %q want %q", store.winCalls[0], player)
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
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertStatusCode(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got: %d\nwant: %d", got, want)
	}
}

func TestLeague(t *testing.T) {
	wantedLeague := []Player{
		{"Chris", 20},
		{"Bob", 30},
		{"Bilbo", 90},
	}
	store := &SpyPlayerStore{
		league: wantedLeague,
	}
	server := NewPlayerServer(store)

	t.Run("returns 200 when get /league", func(t *testing.T) {
		req := makeLeagueRequest()
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		got := getLeagueFromResponse(t, resp.Body)
		assertStatusCode(t, resp.Code, http.StatusOK)

		assertLeague(t, wantedLeague, got)

		assertContentType(t, resp, jsonContentType)
	})
}

func assertContentType(t *testing.T, resp *httptest.ResponseRecorder, contentType string) {
	t.Helper()
	if resp.Result().Header.Get("content-type") != contentType {
		t.Errorf("response did not have content-type of application/json, got %v", resp.Result().Header)
	}
}

func assertLeague(t *testing.T, wantedLeague []Player, got []Player) {
	if !reflect.DeepEqual(wantedLeague, got) {
		t.Errorf("got: %v\nwant: %v\n", wantedLeague, got)
	}
}

func getLeagueFromResponse(t *testing.T, body io.Reader) []Player {
	t.Helper()

	league := []Player{}
	err := json.NewDecoder(body).Decode(&league)
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return league
}

func makeLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}
