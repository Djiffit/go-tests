package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Djiffit/go-tests/time"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := time.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")
	time.NewCLI(store, os.Stdin).PlayPoker()
}
