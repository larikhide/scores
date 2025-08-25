package main

import (
	"log"
	"net/http"
)

type inMemoryStore struct{}

func (s *inMemoryStore) GetPlayerScore(name string) int {
	return 123
}

func (s *inMemoryStore) RecordWin(name string) {
	// TODO: https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server#write-the-test-first-5
}

func main() {
	server := &PlayerServer{
		store: &inMemoryStore{},
	}
	log.Fatal(http.ListenAndServe(":5000", server))
}
