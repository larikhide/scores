package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("could not open the database: %v", err)
	}

	store := &FileSystemPlayerStore{db}
	server := &PlayerServer{
		store: store,
	}

	log.Fatal(http.ListenAndServe(":5000", server))
}
