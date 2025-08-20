package main

import (
	"fmt"
	"net/http"
	"strings"
)

type playerStore interface {
	GetPlayerScore(name string) int
}

type PlayerServer struct {
	store playerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, p.store.GetPlayerScore(player))
}

func getPlayerScore(player string) string {
	if player == "Pepper" {
		return "20"
	}

	if player == "Floyd" {
		return "10"
	}
	return ""
}
