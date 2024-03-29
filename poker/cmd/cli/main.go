package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Djiffit/go-tests/poker"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker my dude")
	fmt.Println("Type {Name} wins to record a win")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
