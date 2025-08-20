package main

import (
	"log"
	"net/http"
)

type inMemoryStore struct{}

func (s *inMemoryStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{
		store: &inMemoryStore{},
	}
	log.Fatal(http.ListenAndServe(":5000", server))
}
