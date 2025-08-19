package main

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	shouldReturn := getPlayerScore(player, w)
	if shouldReturn {
		return
	}
}

func getPlayerScore(player string, w http.ResponseWriter) bool {
	if player == "Pepper" {
		fmt.Fprintf(w, "20")
		return true
	}

	if player == "Floyd" {
		fmt.Fprintf(w, "10")
		return true
	}
	return false
}
