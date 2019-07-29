package main

import (
	"log"
	"net/http"

	"github.com/Djiffit/go-tests/time"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := time.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	server := time.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
