package main

import (
	"log"
	"net/http"

	poker "github.com/tsugoshi/learn-go-application"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()
	server, err := poker.NewPlayerServer(store)

	if err != nil {
		log.Fatalf("error creating PlayerServer %v", err)
	}

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
